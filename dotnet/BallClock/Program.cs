using System;
using System.Collections.Generic;
using System.Diagnostics;

namespace BallClock
{
	class Program
	{
		static void Main(string[] args)
		{
			List<string> input = new List<string>();
			Stopwatch watch = new Stopwatch();
			int numBalls = 0;

			Console.WriteLine("Enter the number of clock balls one per line ranging from 27 to 127 (0 to stop):");

			string line;
			while ((line = Console.ReadLine()) != "0")
			{
				input.Add(line);
			}

			watch.Start();
			foreach (string bc in input)
			{
				if (!int.TryParse(bc, out numBalls) || numBalls < 27 || numBalls > 127)
				{
					Console.WriteLine("'{0}' is not a valid number of balls!", bc);
					continue;
				}

				BallClock clock = new BallClock(numBalls);
				while (true)
				{
					clock.Tick();
					if (clock.AreBallsInSeq)
					{
						Console.WriteLine("{0} balls cycle after {1} days.", clock.NumBalls, clock.Days);
						break;
					}
				}
			}
			Console.WriteLine("Total execution time: {0:ss\\.fff}s", watch.Elapsed);
		}
	}
}
