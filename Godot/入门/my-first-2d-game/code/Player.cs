using Godot;
using System;

public partial class Player : Area2D
{
	[Export]
	public int Speed { get; set; } = 400;
	public Vector2 ScreenSize { get; private set; }

	// 本节点可以 emit “被撞到”信号
	[Signal]
	public delegate void HitEventHandler();

	public override void _Ready()
	{
		ScreenSize = GetViewportRect().Size;

		Hide();
	}

	// 下划线开头的是约定名称的钩子，每帧调用一次
	public override void _Process(double delta)
	{
		var velocity = Vector2.Zero;

		// action 名在项目设置的 input map 里设置好这些名字和对应的按钮
		if (Input.IsActionPressed("move_right"))
			velocity.X += 1;

		if (Input.IsActionPressed("move_left"))
			velocity.X -= 1;

		if (Input.IsActionPressed("move_down"))
			velocity.Y += 1;

		if (Input.IsActionPressed("move_up"))
			velocity.Y -= 1;

		var animatedSprite2D = GetNode<AnimatedSprite2D>("AnimatedSprite2D");

		// 根据是否有速度，决定是否播放动画
		if (velocity.Length() > 0)
		{
			velocity = velocity.Normalized() * Speed; // normalize 一下，避免“斜向速度是 1.414 倍”这种事出现
			animatedSprite2D.Play();
		}
		else
			animatedSprite2D.Stop();

		// 根据速度向量和线速度值，计算下一帧的位置
		Position += velocity * (float)delta;
		Position = new Vector2(
			x: Mathf.Clamp(Position.X, 0, ScreenSize.X),
			y: Mathf.Clamp(Position.Y, 0, ScreenSize.Y)
		);

		// 根据方向不同，播放不同动画。如果动画对称，可以使用 FlipV 和 FlipH 来镜像翻转
		if (velocity.X != 0)
		{
			animatedSprite2D.Animation = "walk";
			animatedSprite2D.FlipV = false;
			animatedSprite2D.FlipH = velocity.X < 0;
		}
		else if (velocity.Y != 0)
		{
			animatedSprite2D.Animation = "up";
			animatedSprite2D.FlipV = velocity.Y > 0;
		}
	}

	private void OnBodyEntered(Node2D body)
	{
		Hide(); // 被撞后，立即隐藏
		EmitSignal(SignalName.Hit);
		// 将本节点下的 CollisionShape2D 设为 Disabled 避免反复触发
		// 注意要使用 SetDeferred，避免在 handler 内直接设置（会出问题）
		GetNode<CollisionShape2D>("CollisionShape2D").SetDeferred(CollisionShape2D.PropertyName.Disabled, true);
	}

	public void Start(Vector2 position)
	{
		Position = position; // 放置到指定位置
		Show(); // 显示出来
		GetNode<CollisionShape2D>("CollisionShape2D").Disabled = false; // 开启碰撞检测
	}
}
