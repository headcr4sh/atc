require 'securerandom'
require 'color'

describe 'dashboard', type: :feature do
  let(:red) { Color::RGB.by_hex('E74C3C') }
  let(:green) { Color::RGB.by_hex('2ECC71') }
  let(:orange) { Color::RGB.by_hex('E67E22') }
  let(:yellow) { Color::RGB.by_hex('F1C40F') }
  let(:brown) { Color::RGB.by_hex('8F4B2D') }
  let(:blue) { Color::RGB.by_hex('3498DB') }
  let(:grey) { Color::RGB.by_hex('ECF0F1') }
  let(:palette) { [red, green, orange, yellow, brown, blue, grey] }

  let(:team_name) { generate_team_name }
  let(:other_team_name) { generate_team_name }

  before(:each) do
    fly_with_input("set-team -n #{team_name} --no-really-i-dont-want-any-auth", 'y')
    fly_with_input("set-team -n #{other_team_name} --no-really-i-dont-want-any-auth", 'y')

    fly_login team_name
    fly('set-pipeline -n -p some-pipeline -c fixtures/states-pipeline.yml')
    fly('unpause-pipeline -p some-pipeline')

    fly_login other_team_name
    fly('set-pipeline -n -p other-pipeline-private -c fixtures/states-pipeline.yml')
    fly('unpause-pipeline -p other-pipeline-private')
    fly('set-pipeline -n -p other-pipeline-public -c fixtures/states-pipeline.yml')
    fly('unpause-pipeline -p other-pipeline-public')
    fly('expose-pipeline -p other-pipeline-public')

    fly_login team_name
  end

  it 'shows all pipelins from the authenticated team and public pipelines from other teams' do
    dash_login team_name

    visit dash_route('/dashboard')

    within '.dashboard-team-group', text: team_name do
      expect(page).to have_content 'some-pipeline'
    end

    within '.dashboard-team-group', text: other_team_name do
      expect(page).to have_content 'other-pipeline-public'
      expect(page).to_not have_content 'other-pipeline-private'
    end
  end

  context 'when a pipeline is paused' do
    before do
      fly('pause-pipeline -p some-pipeline')
    end

    it 'is shown in blue' do
      expect(border_color.closest_match(palette)).to eq(blue)
    end
  end

  context 'when a pipeline has a failed build' do
    before(:each) do
      fly('set-pipeline -n -p some-other-pipeline -c fixtures/states-pipeline.yml')
      fly('unpause-pipeline -p some-other-pipeline')
      fly('set-pipeline -n -p another-pipeline -c fixtures/states-pipeline.yml')
      fly('unpause-pipeline -p another-pipeline')
      fly_fail('trigger-job -w -j some-other-pipeline/failing')
    end

    it 'is shown in red' do
      expect(border_color('some-other-pipeline').closest_match(palette)).to eq(red)
    end

    it 'rises to the top' do
      visit_dashboard
      within '.dashboard-team-group', text: team_name do
        expect(page.find_all('.dashboard-pipeline-name').map(&:text)).to eq ['some-other-pipeline', 'some-pipeline', 'another-pipeline']
      end
    end
  end

  context 'when a pipeline has a passed build' do
    before do
      fly('trigger-job -w -j some-pipeline/passing')
    end

    it 'is shown in green' do
      expect(border_color.closest_match(palette)).to eq(green)
    end
  end

  context 'when a pipeline has an aborted build' do
    before do
      fly('trigger-job -j some-pipeline/running')
      fly('abort-build -j some-pipeline/running -b 1')
    end

    it 'is shown in brown' do
      expect(border_color.closest_match(palette)).to eq(brown)
    end
  end

  context 'when a pipeline is pending' do
    it 'is shown in grey' do
      expect(border_color.closest_match(palette)).to eq(grey)
    end
  end

  context 'when a pipeline has an errored build' do
    before do
      fly_fail('trigger-job -w -j some-pipeline/erroring')
    end

    it 'is shown in amber' do
      expect(border_color.closest_match(palette)).to eq(orange)
    end
  end

  context 'when a pipeline changes its state' do
    it 'updates the dashboard automatically' do
      expect(border_color.closest_match(palette)).to eq(grey)
      fly('trigger-job -j some-pipeline/passing')
      sleep 5
      expect(border_color.closest_match(palette)).to eq(green)
    end
  end

  private

  def login
    @login ||= dash_login team_name
  end

  def border_color(pipeline = 'some-pipeline')
    visit_dashboard
    pipeline = page.find('.dashboard-pipeline', text: pipeline)
    by_rgb(computed_style(pipeline.find('.dashboard-pipeline-banner'), 'backgroundColor'))
  end

  def computed_style(node, attribute)
    page.evaluate_script("window.getComputedStyle(document.evaluate('#{node.path}', document, null, XPathResult.FIRST_ORDERED_NODE_TYPE, null).singleNodeValue).#{attribute}")
  end

  def by_rgb(rgb)
    /rgb\((\d+),\s*(\d+),\s*(\d+)\)/.match(rgb) do |m|
      Color::RGB.new(m[1].to_i, m[2].to_i, m[3].to_i)
    end
  end

  def visit_dashboard
    login
    visit dash_route('/dashboard')
  end
end
