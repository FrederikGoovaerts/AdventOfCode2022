package main

import (
	"fmt"
)

const EX1 = "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
const EX2 = "bvwbjplbgvbhsrlpgdmjqwftvncz"
const EX4 = "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"
const INPUT = "vmsvstvtzvzqvzvssvddsgdddrqrsqqwnnffrwrswwztwzzbssgtstvsshrsrrlvrvzrrsqsvqsqbssqrsrjrsrddqnnfqnnrfnrnffjsjnnqddjjbvvbzzshhsffmtftrrpzznlltclltmltmllzclcwwgzzzqvqhqllcrllqdqpqnqbqggcrcsclssrfsssqhhfzzprrjvrjvvblvbllwccztzggspgggqrqcqjcqjcqcmmffmtmltlwwcpplpqqwhwvhvnnqhqsqlqvvgfgqqddlhhstthzhgzgbzbdzbzlzwlzlpljlgjjqlltrrbmrrvvcttsddfjddmbdbffcrfftvvzppncpcvcdvcvvjjjlttcgtcctrrrgmrgmgwwsgscsffswfwqwnnzmzffpwptpntppbddggsbsddmjmssnzsnstsztttjztjjpddwlwrrlnnzsnnfllqrrqmqhhfddtmmpfmfcfgfllngnvnjjmffwbffsqsmsdstdddjssmlssfpprrgzzzwvvftffczfftccwrrdsrsjrssfnsnttqptqtnqtnqtttjqtjqjzjpzpvzvbzvvjqqvcqcwwgwllbwbjjhsjhhchffgdfgffvddrndrdwdwlwqqfqzzntzzhwzzndnjnsjjmvvlqqqqshhmzzjvvcscqqjjhphnppcprrhnnbsspcscbssnsmmznzsnnszsrzrwrswsgwglwlzznpnwpnnjlnlrnntllhvvdrvvprpptqqlzqzrrgfgfhfvhhcscjjfmmjpmjjvwvtvhvqhqvvqvwvpvsschhlmlrlmljjqnqggbwbmwbwgbwwgllqzllrjlrjjfjzffthhpvvsddvjvvzsvvhghlhqqbnqqlnngsngnccnpccmttmrtmmlmccqbbjssbsqbqddnqqdvdvbvgbvvjfjhjzzfhfzffgccdjcjtjpttcrrffvddwgglrgllbttcqtctqtvtwtggrppzvvbzvzgzsznnsqnsqqnwwzrwwnwbwnnhznhzhchwwvwhwlwlnnmwmfmfrmmnvndnffjgjmgmfmhmtmstthjhbjhjddrqrllnflnnwzwjjthjhwhmmflmlfmfddqgdqgqqwbbdsstssrvvbpbgppfwfvvglvvzgvgmvgvzvdvcctddjqddffrprbrggpgpghhwzhzggqrqsrstrssnttrgrsswsrstsdtdzttzvznzjnzzqssjwswcswsgghqghqgqvggtsswwmbbfrrdcctjtjsttrccqbbtcctcscqsssvmmdhmdmvdmvmsvsrvsrvsszvvrpvrpprvppctclcrrchhtmmbddswddglgrrdcdppcbpbgpbggbssnrnbrbllqwqbwbmmcwmmrjrdjjpsjppblbjbsbsgswsmwmgwgmmqpplzlssvcscnngwgvvffvlvtvtddgqgttjljggfzggmfgfqqtrrdbdbbnjnnbjbwwjjqljlnjjvzvhvlhhhsfhsfscsjcjsjbsjjdbdbpbqqsggcnczcchrhjhphlppnwwhthghhfvfrrbzrbrrcsszfmhrjflswthfrlmjbnhblmsldjpnrdgfbcrftnbcltlctbdthhjzsqvbmppcbhctqbtfhdrtrbzjpnzzvtfjtpqgbtmdcnjgtblhmvrpbrcpqfhpzqvbfqhsrjqwqqnggrlgvwndqrhvpzhswnglngwjdgwcnngrvhtcsblpdshqcwfpzcpmrbzqjfpllbbvlcfjtdqchjsgsqgqzsbfnnqcsvpwfmzdhzmcjfjnczldclvpfpntppqbhfqzqdtfwsvdzhphwpzhqqfflswhtldlrmrqpfdlwwszphrwjqvhpqmrzvblvtvsfpgzdhnlmsgpqnqdtbjqccbtnvqmtdggdsvvfwjzbdpnztmsblpgcrmppblmtfnjvjwmhnwqdmdrvbvsvhgbhjwzjfcqzwtvzlfmpzvpwsdsqphmqlzwfrqdjdbsmfdzllqqfhmcmhchjpbqtjbczhbcllmtqgczdfnjsjhhcsfdhwwbzzjdqtlvgfgfwbqztqftfqhsfvbzjcbmtdzszfsgzpqpvtqnmjhlrtbvfphvhwbnqsnwgpznnwcbbvdqvvzllqpjbqhqwtfzlhqgvmrpjmlvczqgmffsjvgzvnvccgnmdmbcqjqvbnzvdgnbdwbszldjwjdsnplgqnjmlvrlzcvlgtrdndhjmrtczqjqrzzzclrdmdmrgvljstqfhldfzgvhcdtmncqjjnghgwdgnhbjztfgpqnvfhzngqcftpbqsnvgrqpczwptghbssljnftzmbrqmnbvcbsshvmmmczjgnbjwltflndtntmfznmqhzjjtmmtwbflhmqcmlmzrnvfqzmzhhszqmmqwgzcnncjwvdzvczsgbpcscjptfpwhvvvstvqzhtptrmbhttcrdzhljpmclrbnzmvmtbmmdhfgpjjwsldsgdjflhmfrgvmgzvntgmlglfsntpjjwgnpwsjvsnbctwlsqvprlfmsqfdmzdsdvbtsnqzflvtlgrmdhfgfmwrcqjvbwqsblnwlmtqfhdwwlchqljbljwzmmqgqpwfnmjlvhmppzsdjlbvqwmwqwqqsrgzmpzzzzstrtrbwttzjcmlssstpntnrwtgthrfmthbjjmmtcvjnwsdpdndccgncjbjdqbwtlbfcmgvmlbrdzzlzcpctjvldbpvvmsrwrhhtdsrwljtmntcftwvppjvtttgslrdvpglzdnhrgnrzjsczsmtcfclhvncrhsfvbppqrbwsmjbhjmllwwlqvhbljrqpzshqvllzqtjfhdlnghwblzjldcmfgfjgcbfwqldvdppjmsqqqmznwtcgpmlpqmdpzctwjpnstvgjsbzspfbvfcrgzzcdbzfthgbjrnjwpprblnhdcfgpdcmljpqqzchslpvlwqjvsqtnvnfjhpvwjnjzghnbslhmrhwdppgtvzzmhzzzsvldphjlhrfjwrctfcmnzpvvwqrczcsmznflcwzsplrsdvmfghmnbbjdzclnzpbtjllcszzjcqbbczqhbnfpcctbmhfglsdtnzdjrjbjqczqflznfwzsrmpcgvgrjmhtpgllcpdnlrzvqljprnpvcglvmbtbzjwqfrhdngwsrfjsqnhfncnvprrcngjzfdvcnphcchbqhqpbbbzgmbfndlffnrpzvplptnrvjlzfczvsbctmdqrzdnpvrtgnsgdjwqshglspdzbhflfbsfvbrqgfzmhzhshdlcnqhzbnhbhmsgffgsmvglhscqsrsvmszjfgdhsglbqgwwjndscsqvpccwhlvjbtvftdppcscjblwtmpvvchpzwmblmbbrrvwlglfqwtgcffnzljnmtpbcrszvblwfqdslpqlrmnvrhvrzwnprnmntzrpsdcnwfgljldpvjwwzzbpcltnvghqvqnvmmrfqsnbstdctqqgtpzqttnrrdstrtlfzmvvbzzwwchrscqlnpzdmphdwdqdwbszlhwsbfscthndrdvhtgbpsmznlfjpwjchvzbmrflcphhgjfwstjnzlllztgzjmnwcglmhrbztzdnbcbrgrlmpprrbthbhslfrsjsrrrqvwmqghcgdvvsmrqdwwbdnzmnsqfflpbbjvzjdgsfjdjmbhptmrnbqpqhcdwtvbdlrdzsrchqlsrccjfbfrnnrctdsqbnjzzvrlcwphsscppdmsbrqrzddrdvjwfrldccmzwrhrbpcqpnvnbgpsrhvbchmrdtwqrlvpwhqgnppmsdvqdldlrjrzntlsfwmwhjsghddsppchqlltrhlwrccvflwjbvptclqvwdmrmghbngbflfspsjmglzcqjdtdtldmmngljwfqvfwnmfdtrsmlqzszhzccgvbnnwtpbgssrcqlqgrsjqfhhwpvdtsclgsntwvcsfpvwfqnmwhmtldfswqrsvzshfzwlnwhqhwsrqqzlsdhvqfwnhjcvmplcljmlhbtqtrpjfjfdfmlhtfpnszptfnjbldscjgvjhpzflhm"

func solve(input string, messageLength int) int {
	for i := 0; i < len(input); i++ {
		duplicate := false
		for start := i; start < i+messageLength; start++ {
			for check := start + 1; check < i+messageLength; check++ {
				if input[check] == input[start] {
					duplicate = true
				}
			}
		}
		if !duplicate {
			return i + messageLength
		}
	}
	return -1
}

func main() {
	input := INPUT

	fmt.Println(solve(input, 4))
	fmt.Println(solve(input, 14))
}
