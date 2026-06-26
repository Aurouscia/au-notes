using Godot;

public partial class Mob : RigidBody2D
{
	public override void _Ready()
	{
		var animatedSprite2D = GetNode<AnimatedSprite2D>("AnimatedSprite2D");
		string[] mobTypes = animatedSprite2D.SpriteFrames.GetAnimationNames();
		// 使得 AnimatedSprite2D 节点随机选一种动画开始播放
		animatedSprite2D.Play(mobTypes[GD.Randi() % mobTypes.Length]);
	}

	public override void _Process(double delta)
	{
	}

	// 让 VisibleOnScreenNotifier2D 节点检测到本 mob 离开屏幕时，移除本 mob
	private void OnVisibleOnScreenNotifier2DScreenExited()
	{
		QueueFree();
	}
}
