using Godot;
using System;

public partial class Mob : CharacterBody3D
{
	[Export]
	public int MinSpeed { get; set; } = 10;
	[Export]
	public int MaxSpeed { get; set; } = 18;
	[Signal]
	public delegate void SquashedEventHandler();

	public override void _PhysicsProcess(double delta)
	{
		MoveAndSlide();
	}

	public void Initialize(Vector3 startPosition, Vector3 playerPosition)
	{
		// 将玩家位置投影到与出生点同一高度，避免玩家跳跃时 mob 抬头
		playerPosition.Y = startPosition.Y;
		// 将节点移动到指定位置，然后看向指定位置
		LookAtFromPosition(startPosition, playerPosition, Vector3.Up);
		// 随机旋转 -45～45 度
		RotateY((float)GD.RandRange(-Mathf.Pi/4, Mathf.Pi/4));

		// 根据速度上下限随机一个速度值
		int randomSpeed = GD.RandRange(MinSpeed, MaxSpeed);
		// 根据速度调整动画播放速度
		GetNode<AnimationPlayer>("AnimationPlayer").SpeedScale = randomSpeed / MinSpeed;
		// 根据速度标量创建速度向量
		Velocity = Vector3.Forward * randomSpeed;
		// Velocity 是全局的，所以要朝着 Mob 的“前方”的话需要手动旋转
		// 根据当前的朝向（Rotation.Y）旋转速度向量，沿着 Up 轴
		Velocity = Velocity.Rotated(Vector3.Up, Rotation.Y);
	}

	// 将 Notifier 的“离开屏幕”信号连接到 Mob 节点的本函数上
	public void OnVisibilityNotifierScreenExited()
	{
		QueueFree();
	}

	public void Squash()
	{
		Console.WriteLine("Squash");
		EmitSignal(SignalName.Squashed);
		QueueFree();
	}
}
