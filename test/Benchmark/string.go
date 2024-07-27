package benchmark
import(
	"strings"
)
func ConcatStringWithoutBuilder(a,b string) string {
	var s string=""
    for i:=0;i<1000;i++ {
		s+=a+b
	}
	return s
}
func ConcatStringWithBuilder(a,b string) string {
	var sb strings.Builder
    for i:=0;i<1000;i++ {
		sb.WriteString(a+b)
	}
	return sb.String()
}