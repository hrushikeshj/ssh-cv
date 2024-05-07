package cv

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) renderEduHeader(uni, date, degree, cgpa string) string {
	buf := 5
	middle_buf := m.width - len(uni) - buf
	var head string = lipgloss.JoinHorizontal(lipgloss.Bottom,
		m.styles.subSectionHeader.Render(uni),
		m.styles.duration.Width(middle_buf).Align(lipgloss.Right).Render(date))

	middle_buf = m.width - len(degree) - buf
	var degree_head string = lipgloss.JoinHorizontal(lipgloss.Bottom,
		m.styles.subSectionHeader.Render(degree),
		m.styles.duration.Copy().Width(middle_buf).Align(lipgloss.Right).Italic(true).Render("CGPA: "+cgpa))

	return lipgloss.JoinVertical(lipgloss.Left, head, m.styles.companyDisplay.Render(degree_head))

}

func (m Model) renderProjectHeader(name, tech string) string {
	var div string = lipgloss.NewStyle().Foreground(special).Render(" | ")

	return lipgloss.JoinHorizontal(lipgloss.Bottom, name, div, m.styles.techText.Render(tech))
}

// Render role Header
func (m Model) renderRole(role, date, company string) string {
	buf := 5
	middle_buf := m.width - len(role) - buf
	var head string = lipgloss.JoinHorizontal(lipgloss.Bottom,
		m.styles.subSectionHeader.Render(role),
		m.styles.duration.Width(middle_buf).Align(lipgloss.Right).Render(date))

	return lipgloss.JoinVertical(lipgloss.Left, head, m.styles.companyDisplay.Render(company))

}

func (m Model) renderPoints(points []string) string {
	newPoints := []string{}
	for i := 0; i < len(points); i++ {
		s := lipgloss.JoinHorizontal(lipgloss.Top, "• ", m.styles.bulletWidth.Render(points[i]))
		newPoints = append(newPoints, s)
	}

	return lipgloss.JoinVertical(lipgloss.Left, newPoints...)
}

func (m Model) renderBlock(roleHead string, points []string) string {
	return lipgloss.JoinVertical(lipgloss.Left, roleHead, m.styles.bulletBlock.Render(m.renderPoints(points)))
}

func (m Model) RenderCV() string {
	doc := strings.Builder{}

	doc.WriteString(m.styles.nameTitle.Render("\nHrushikesh J") + "\n\n")
	doc.WriteString(m.styles.sectionHeader("Education") + "\n")

	for _, edu := range MyCV.education {
		edu_block := m.renderEduHeader(edu.uni, edu.start+" - "+edu.end, edu.degree, edu.cgpa)
		doc.WriteString(m.styles.sectionBlock(edu_block))

		doc.WriteString("\n\n")
	}

	doc.WriteString(m.styles.sectionHeader("Experience") + "\n")
	for _, experience := range MyCV.experiences {
		doc.WriteString(
			m.styles.sectionBlock(m.renderBlock(m.renderRole(experience.role, experience.start+" - "+experience.end, experience.company),
				experience.points)))

		doc.WriteString("\n\n")
	}

	doc.WriteString(m.styles.sectionHeader("Projects") + "\n")
	for _, project := range MyCV.projects {
		project_block := m.renderBlock(m.renderProjectHeader(project.name, project.tech), project.points)
		doc.WriteString(m.styles.sectionBlock(project_block))

		doc.WriteString("\n\n")
	}

	doc.WriteString(m.styles.sectionHeader("Technical Skills") + "\n")
	skl := strings.Builder{}
	for _, skill := range MyCV.skills {
		skl.WriteString(lipgloss.NewStyle().Faint(true).Render(skill.typ+": ") + skill.skills)
		skl.WriteRune('\n')
	}
	doc.WriteString(m.styles.sectionBlock(skl.String()))

	// doc.WriteString(
	//     sectionBlock(renderBlock(renderRole("Web Developer", "March 2021 - Present", "IRIS, NITK"),
	//     []string{"MarginLginLef", "MarginLef MarginLefginLef Margi nLef Mar ginLef Margi nLef ginLef MarginLef Mar ginLef Marg inLefginLef MarginLef"})))

	return m.styles.body.Render(doc.String())
}
