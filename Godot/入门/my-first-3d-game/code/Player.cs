using Godot;
using System;

public partial class Player : CharacterBody3D
{
	[Export]
	public int Speed { get; set; } = 14;

	[Export]
	public int FallAcceleration { get; set; } = 75;

	private Vector3 _targetVelocity = Vector3.Zero;

	[Export]
	public float TurnSpeed { get; set; } = 10.0f;

	[Export]
	public float JumpImpulse { get; set; } = 20;

	[Export]
	public int BounceImpulse { get; set; } = 16;

	// Emitted when the player was hit by a mob.
	[Signal]
	public delegate void HitEventHandler();

	private AnimationPlayer _animationPlayer;
	private Node3D _pivot;

  	public override void _Ready()
  	{
		_animationPlayer = GetNode<AnimationPlayer>("AnimationPlayer");
		_pivot = GetNode<Node3D>("Pivot");
  	}

	public override void _PhysicsProcess(double delta)
	{
		Vector3 direction = Vector3.Zero;

		// 3D中，“地面”是 XZ 平面

		if (Input.IsActionPressed("move_left"))
			direction.X -= 1.0f;
		if (Input.IsActionPressed("move_right"))
			direction.X += 1.0f;
		if (Input.IsActionPressed("move_forward"))
			direction.Z -= 1.0f;
		if (Input.IsActionPressed("move_back"))
			direction.Z += 1.0f;
		
		if (direction != Vector3.Zero)
		{
			// 单位化，避免斜向运动时产生 1.414 的长度
			direction = direction.Normalized();
			
			// （直接转向）设置 Pivot 的局部坐标系，使其 -Z 方向朝向 direction，同时 Y 轴尽量朝上。如果 direction 不是单位向量，就会错误地引入缩放。
			// GetNode<Node3D>("Pivot").Basis = Basis.LookingAt(direction);

			// （插值转向）获取原本的朝向，然后插值旋转过去
			var pivot = GetNode<Node3D>("Pivot");
			var currentDirection = -pivot.Basis.Z;
			var newDirection = currentDirection.Slerp(direction, (float)delta * TurnSpeed).Normalized(); // Slerp 的第二参数是“步长”而非权重
			pivot.Basis = Basis.LookingAt(newDirection);

			// 4 倍速度播放动画
			_animationPlayer.SpeedScale = 4;
		}
		else
		{
			// 原速度播放动画
			_animationPlayer.SpeedScale = 1;
		}

		_targetVelocity.X = direction.X * Speed;
		_targetVelocity.Z = direction.Z * Speed;

		// IsOnFloor 内置方法可以判断当前角色脚下是否有“地板”。如果脚下没有地板，则让其“往下掉”
		if (!IsOnFloor())
		{
			_targetVelocity.Y -= FallAcceleration * (float)delta;
		}
		else
		{
			if (Input.IsActionJustPressed("jump"))
			{
				_targetVelocity.Y = JumpImpulse;
			}
		}

		// Iterate through all collisions that occurred this frame.
		for (int index = 0; index < GetSlideCollisionCount(); index++)
		{
			// We get one of the collisions with the player.
			KinematicCollision3D collision = GetSlideCollision(index);

			// If the collision is with a mob.
			// With C# we leverage typing and pattern-matching
			// instead of checking for the group we created.
			if (collision.GetCollider() is Mob mob)
			{
				// We check that we are hitting it from above.
				if (Vector3.Up.Dot(collision.GetNormal()) > 0.1f)
				{
					// If so, we squash it and bounce.
					mob.Squash();
					_targetVelocity.Y = BounceImpulse;
					// Prevent further duplicate calls.
					break;
				}
			}
		}

		// 计算完成，指定物体的速度为该速度（让其能实际运动）
		Velocity = _targetVelocity; 
		MoveAndSlide();

		_pivot.Rotation = new Vector3(Mathf.Pi / 6.0f * Velocity.Y / JumpImpulse, _pivot.Rotation.Y, _pivot.Rotation.Z);
	}

	private void Die()
	{
		EmitSignal(SignalName.Hit);
		QueueFree();
	}

	// We also specified this function name in PascalCase in the editor's connection window.
	private void OnMobDetectorBodyEntered(Node3D body)
	{
		Die();
	}
}
