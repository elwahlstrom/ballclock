using System;
using System.Collections.Generic;

namespace BallClock
{
	public class BallClock
	{
		//members
		private Stack<int> _1minTrack;
		private Stack<int> _5minTrack;
		private Stack<int> _1hourTrack;
		private Queue<int> _queue;

		#region Constructor

		/// <summary>
		/// Constructor
		/// </summary>
		/// <param name="numBalls"># of balls</param>
		public BallClock(int numBalls)
		{
			//init 
			this.NumBalls = numBalls;

			//setup tracks
			_1minTrack = new Stack<int>(4);
			_5minTrack = new Stack<int>(11);
			_1hourTrack = new Stack<int>(11);
			_queue = new Queue<int>(numBalls);

			//setup the queue
			for (int i = 1; i <= numBalls; i++)
				_queue.Enqueue(i);
		}

		#endregion

		#region Properties

		/// <summary>
		/// # of balls allocated to the clock
		/// </summary>
		public int NumBalls { get; private set; }

		/// <summary>
		/// # of days
		/// </summary>
		public float Days { get; private set; }

		/// <summary>
		/// # of hours
		/// </summary>
		public int Hours
		{
			get
			{
				return _1hourTrack.Count;
			}
		}

		/// <summary>
		/// # of  minutes
		/// </summary>
		public int Minutes
		{
			get
			{
				return _5minTrack.Count * 5 + _1minTrack.Count;
			}
		}

		/// <summary>
		/// Determines if the balls are in sequence
		/// </summary>
		public bool AreBallsInSeq { get; set; }

		#endregion

		#region Methods

		/// <summary>
		/// Cycle a clock ball
		/// </summary>
		public void Tick()
		{
			int ball = _queue.Dequeue();

			if (_1minTrack.Count < 4)
				_1minTrack.Push(ball);
			else
			{
				EmptyTrack(_1minTrack);
				if (_5minTrack.Count < 11)
					_5minTrack.Push(ball);
				else
				{
					EmptyTrack(_5minTrack);
					if (_1hourTrack.Count < 11)
						_1hourTrack.Push(ball);
					else
					{
						EmptyTrack(_1hourTrack);
						_queue.Enqueue(ball);

						this.Days += .5f;
						this.AreBallsInSeq = CheckSeqOrder();
					}
				}
			}
		}

		/// <summary>
		/// Print the clock
		/// </summary>
		public void Print()
		{
			Console.WriteLine("|-");
			Console.WriteLine("|{0}", ReverseArray(_1minTrack.ToArray()));
			Console.WriteLine("|{0}", ReverseArray(_5minTrack.ToArray()));
			Console.WriteLine("|{0}", ReverseArray(_1hourTrack.ToArray()));
			Console.WriteLine("|-");
			Console.WriteLine("|{0}", this);
			Console.WriteLine("|-");
		}

		public override string ToString()
		{
			return string.Format("{0} days, {1} hours and {2} minutes", this.Days, this.Hours, this.Minutes);
		}

		#endregion

		#region Private Methods

		private void EmptyTrack(Stack<int> track)
		{
			while (track.Count > 0)
			{
				_queue.Enqueue(track.Pop());
			}
		}

		private string ReverseArray(int[] a)
		{
			String s = "";
			for (int i = a.Length - 1; i >= 0; i--)
				s += String.Format("{0},", a[i]);
			return s.TrimEnd(',');
		}

		private bool CheckSeqOrder()
		{
			int i = 1;
			foreach (int ball in _queue)
			{
				if (i != ball)
					return false;
				i++;
			}
			return true;
		}

		#endregion
	}
}
