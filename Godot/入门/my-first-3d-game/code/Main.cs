using Godot;
using System;

public partial class Main : Node
{
	[Export]
	public PackedScene MobScene { get;set; }

	private Control _retry;

  	public override void _Ready()
  	{
		// 游戏启动时，初始隐藏 retry 蒙层
		_retry = GetNode<Control>("UserInterface/Retry");
		_retry?.Hide();	
  	}

	public void OnMobTimerTimeout()
	{
		// 实例化 mob 对象
		var mob = MobScene.Instantiate<Mob>();

		// 获取 SpawnLocation 和 Player 对象（注意如果有多层，要使用斜杠分隔）
		var spawnLocation = GetNode<PathFollow3D>("SpawnPath/SpawnLocation");
		var player = GetNode<Player>("Player");

		// 随机一个 0～1 的浮点数，用于设置 Path 上的“比例”
		spawnLocation.ProgressRatio = GD.Randf();

		// 以当前玩家的位置和 SpawnLocation 初始化 mob
		mob.Initialize(spawnLocation.Position, player.Position);
		
		// 将 mob 添加到根节点下
		AddChild(mob);

		// We connect the mob to the score label to update the score upon squashing one.
		mob.Squashed += GetNode<ScoreLabel>("UserInterface/ScoreLabel").OnMobSquashed;
	}

	private void OnPlayerHit()
	{
		// 显示 retry 蒙层
		_retry?.Show();

		// 停止 mob 生成定时器
		GetNode<Timer>("MobTimer").Stop();
	}

  	public override void _UnhandledInput(InputEvent @event)
	{
		// 如果当前按下了 Enter 键，而且 retry 蒙层可见（当前玩家是“已死亡”状态）
		if(@event.IsActionPressed("retry") && _retry.Visible)
		{
			GetTree().ReloadCurrentScene(); // 重新开始整个 main 场景
		}
	}
}
