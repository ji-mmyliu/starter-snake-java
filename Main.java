import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;

public class Main {
    // Welcome to Battlesnake!
    // This is a Java Battlesnake server that implements a simple input/output
    // system described below

    public static void main(String[] args) throws IOException {
        char grid[][] = new char[11][11]; // The battlesnake grid will be 11x11

        // use this Buffered Reader to read input
        BufferedReader br = new BufferedReader(new FileReader("input.txt"));

        /**
         * TODO: Read the 11x11 char array as the Battlesnake board
         * (an example of the input can be found in sample_input.txt)
         * 
         * "H" represents your snake's head
         * "B" represents a segment of your snake's body
         * "T" represents your snake's tail
         * 
         * "h" represents an enemy snake's head
         * "b" represents a segment of an enemy snake's body
         * "t" represents a an enemy snake's tail
         * 
         * "." represents empty space
         * "f" represents a food pellet
         */
        for (int i = 0; i < 11; i++) {
            String line = br.readLine().trim();
            grid[i] = line.toCharArray();
        }

        /**
         * This is the remaining health of your snake (a number from 1 to 100)
         */
        int health_remaining = Integer.parseInt(br.readLine());

        /**
         * TODO: output the move that YOUR snake will make (either "up", "down", "left",
         * or "right")
         */
        System.out.println("up"); // just an example. replace this with the move of your choice
    }
}