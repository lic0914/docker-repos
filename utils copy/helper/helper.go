package helper

import(
	"strings"
)

type StringBuilder strings.Builder

func(sb *StringBuilder) Append(s string) StringBuilder{
	sb.WriteString(s)
	return sb
}
func(sb *StringBuilder) AppendLine() StringBuilder{
	sb.Append("\n")
	return sb
}
