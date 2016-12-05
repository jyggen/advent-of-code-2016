package main

import (
	"github.com/jyggen/advent-of-go/util"
	"strconv"
)

func solve(input string) (int, int) {
	/*// He begins by delivering a present to the house at his starting location.
	  $houses      = [md5('0x0') => null];
	  $inputLength = count($input);

	  for ($santa = 0; $santa < $numberOfSantas; $santa++) {
	      // Santa is delivering presents to an infinite two-dimensional grid of houses.
	      $xAxis = 0;
	      $yAxis = 0;

	      for ($step = $santa; $step < $inputLength; $step += $numberOfSantas) {
	          // Moves are always exactly one house to the north (^), south (v), east (>), or west (<).
	          switch ($input[$step]) {
	              case '^':
	                  $yAxis--;
	                  break;
	              case 'v':
	                  $yAxis++;
	                  break;
	              case '>':
	                  $xAxis++;
	                  break;
	              case '<':
	                  $xAxis--;
	                  break;
	          }

	          $houseId = md5($xAxis.'x'.$yAxis);

	          // After each move, he delivers another present to the house at his new location.
	          if (isset($houses[$houseId]) === false) {
	              $houses[$houseId] = null;
	          }
	      }
	  }

	  return count($houses);
	*/
	return 0, 0
}

func main() {
	part1, part2 := solve(util.ReadFile("2015/03/input"))

	util.PrintAnswers(strconv.Itoa(part1), strconv.Itoa(part2))
}
