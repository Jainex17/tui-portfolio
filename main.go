package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	index  int
	width  int
	height int
}

func New() *model {
	return &model{index: 1}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "right":
			if m.index == 3 || m.index == 0 {
				return m, nil
			}
			m.index++
		case "left":
			if m.index == 1 || m.index == 0 {
				return m, nil
			}
			m.index--
		case "esc", "q":
			return m, tea.Quit
		}
	}

	return m, cmd
}

func (m model) View() string {
	if m.width == 0 {
		return "loading..."
	}
	navbox := lipgloss.NewStyle().
		Width(m.width - 50).
		PaddingTop(3).
		Align(lipgloss.Left)

	navlink1 := `Home`
	navlink2 := `About`
	navlink3 := `Projects`

	navstyle := lipgloss.NewStyle().
		Width(10).
		Align(lipgloss.Center)

	navstyleActive := navstyle.Copy().
		Background(lipgloss.Color("5"))

	leftnavStyle := lipgloss.NewStyle().
		Width(30).
		Align(lipgloss.Left)

	navLeft := lipgloss.JoinVertical(lipgloss.Center,
		leftnavStyle.Render("Portfolio | jainex.xyz"),
	)

	nav := lipgloss.JoinHorizontal(lipgloss.Right,
		lipgloss.NewStyle().Render(func() string {
			if m.index == 1 {
				return navstyleActive.Render(navlink1)
			}
			return navstyle.Render(navlink1)
		}()),
		lipgloss.NewStyle().Render(func() string {
			if m.index == 2 {
				return navstyleActive.Render(navlink2)
			}
			return navstyle.Render(navlink2)
		}()),
		lipgloss.NewStyle().Render(func() string {
			if m.index == 3 {
				return navstyleActive.Render(navlink3)
			}
			return navstyle.Render(navlink3)
		}()),
	)

	navboxWidth := m.width - 50
	gap := navboxWidth - lipgloss.Width(navLeft) - lipgloss.Width(nav)
	if gap < 1 {
		gap = 1
	}
	spacer := lipgloss.NewStyle().Width(gap).Render("")

	navRendered := navbox.Render(
		lipgloss.JoinHorizontal(lipgloss.Bottom, navLeft, spacer, nav),
	)

	footerLeftText := "press <- or -> to navigate"
	footerRightText := "press 'q' or 'esc' to quit"

	footerLeft := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Render(footerLeftText)

	footerRight := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Render(footerRightText)

	footerGap := m.width - lipgloss.Width(footerLeftText) - lipgloss.Width(footerRightText) - 50
	if footerGap < 1 {
		footerGap = 1
	}
	footerCenterGap := lipgloss.NewStyle().Width(footerGap).Render("")

	footerRendered := lipgloss.JoinHorizontal(lipgloss.Bottom, footerLeft, footerCenterGap, footerRight)

	content := `
		██╗  █████╗  ██╗ ███╗  ██╗ ███████╗ ██╗  ██╗
		██║ ██╔══██╗ ██║ ████╗ ██║ ██╔════╝ ╚██╗██╔╝
		██║ ███████║ ██║ ██╔██╗██║ █████╗    ╚███╔╝ 
   ██╗  ██║ ██╔══██║ ██║ ██║╚████║ ██╔══╝    ██╔██╗ 
   ╚█████╔╝ ██║  ██║ ██║ ██║ ╚███║ ███████╗ ██╔╝╚██╗
	╚════╝  ╚═╝  ╚═╝ ╚═╝ ╚═╝  ╚══╝ ╚══════╝ ╚═╝  ╚═╝
   `

	switch m.index {
	case 2:
		contentTop := `$ cat about.txt`
		
		aboutText := `A Full Stack Developer, passionate about building things on the internet. 
I work with Frontend, Backend, and App development - whatever gets the 
job done. Currently working as a full-stack developer. Feel free to 
reach out for any exciting projects or crazy ideas.`

		skillsHeader := `$ cat skills.txt`

		languagesLine := lipgloss.NewStyle().Foreground(lipgloss.Color("150")).Underline(true).Render("Languages:") + " JavaScript, TypeScript, Python, Java, SQL"
		webDevLine := lipgloss.NewStyle().Foreground(lipgloss.Color("150")).Underline(true).Render("Web & App Dev:") + " React.js, Next.js, Vue.js, Tailwind CSS, Redux, React Native"
		backendLine := lipgloss.NewStyle().Foreground(lipgloss.Color("150")).Underline(true).Render("Backend Dev:") + " Node.js, Express.js, MongoDB, PostgreSQL, Prisma, Firebase"
		toolsLine := lipgloss.NewStyle().Foreground(lipgloss.Color("150")).Underline(true).Render("Tools & Tech:") + " Git, Docker, VS Code, Cursor"
		
		skillsContent := lipgloss.JoinVertical(lipgloss.Left,
			"",
			languagesLine,
			"",
			webDevLine,
			"",
			backendLine,
			"",
			toolsLine,
		)

		contactHeader := `$ cat contact.txt`

		contactContent := `
Email: jainexp017@gmail.com
Website: jainex.xyz
LinkedIn: linkedin.com/in/jainex17
GitHub: github.com/jainex17`

		content = lipgloss.JoinVertical(lipgloss.Left, 
			lipgloss.NewStyle().Foreground(lipgloss.Color("202")).Render(contentTop),
			lipgloss.NewStyle().MarginTop(1).Render(aboutText),
			lipgloss.NewStyle().Foreground(lipgloss.Color("202")).MarginTop(2).Render(skillsHeader),
			skillsContent,
			lipgloss.NewStyle().Foreground(lipgloss.Color("202")).MarginTop(2).Render(contactHeader),
			lipgloss.NewStyle().Render(contactContent),
		)
	
	case 3:
		projectsTop := `$ cat projects.txt`
		
		webAppsLine := lipgloss.NewStyle().Foreground(lipgloss.Color("150")).Underline(true).Render("Web Applications")
		webAppsContent := `
blaze - Create website in seconds (3k+ visitors, 400+ signups, 500+ website generated)
GrabPost - Grab amazing images for your content
CoSnippet - Platform to share code snippets
RepoVerifier - Verify GitHub repo originality
Coinplay - Gambling simulation :)
PicShareX - Anonymous image sharing platform
CheerMe - Support platform with Stripe integration`

		blockchainLine := lipgloss.NewStyle().Foreground(lipgloss.Color("150")).Underline(true).Render("Blockchain & Web3")
		blockchainContent := `
ETHCinemaNation - Ethereum-based movie rating platform`

		mobileLine := lipgloss.NewStyle().Foreground(lipgloss.Color("150")).Underline(true).Render("Mobile Application")
		mobileContent := `
WallSpace - Get Unique Wallpapers for you mobile`

		extensionsLine := lipgloss.NewStyle().Foreground(lipgloss.Color("150")).Underline(true).Render("Extensions & Bot")
		extensionsContent := `
create-ex-ts - Instant Express + TypeScript setup (NPM Package)
Calendra - amazing calendar for your React app (NPM Package)
AniQuiz - Anime Quiz and AI ChatBot (Discord Bot)`

		projectsContent := lipgloss.JoinVertical(lipgloss.Left,
			"",
			webAppsLine,
			webAppsContent,
			"",
			blockchainLine,
			blockchainContent,
			"",
			mobileLine,
			mobileContent,
			"",
			extensionsLine,
			extensionsContent,
		)

		content = lipgloss.JoinVertical(lipgloss.Left, 
			lipgloss.NewStyle().Foreground(lipgloss.Color("202")).Render(projectsTop),
			projectsContent,
		)
	}

	contentHeight := m.height
	if m.index >= 1 {
		contentHeight = m.height - 20
	}

	box := lipgloss.NewStyle().
		Width(m.width).
		Height(contentHeight).
		Align(lipgloss.Center, lipgloss.Center)

	contentStyle := lipgloss.NewStyle().
		Width(m.width-50).
		Height(m.height-10).
		Padding(0).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("5"))

	if m.index == 1 {
		contentStyle = contentStyle.Align(lipgloss.Center, lipgloss.Center).PaddingRight(4)
	} else {
		contentStyle = contentStyle.Padding(1, 3, 1, 3)
	}

	contentRendered := contentStyle.Render(content)
	boxRendered := box.Render(contentRendered)

	return lipgloss.JoinVertical(lipgloss.Center, navRendered, boxRendered, footerRendered)

}

func main() {

	m := New()

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
