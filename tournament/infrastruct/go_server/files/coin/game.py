
# Receive tho files as arguments in cli
import sys
import random
import os

if __name__ == "__main__":
    # Check if there are 3 arguments in total, self + 2 files
    if len(sys.argv) != 3:
        print("Amount of arguments is not correct, given {}".format(len(sys.argv)))
        sys.exit(1)
    
    winner = random.randint(1,3) # p1 wins, p2 wins, tie -> 1, 2, 3
    print(winner)
        