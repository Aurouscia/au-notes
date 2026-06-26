using Godot;

public partial class Main : Node
{
	// 记得重新 build 项目，否则编辑器里看不到新的属性

	// C# 里没有 preload，动态加载场景只能使用 var scene = ResourceLoader.Load<PackedScene>("res://scene.tscn")
	[Export]
	public PackedScene MobScene { get; set; }

	private int _score;

	public override void _Ready()
	{
		// NewGame(); // 注释掉，否则会自动开始游戏
	}

	public void GameOver()
	{
		GetNode<AudioStreamPlayer>("Music").Stop(); // 停止播放 BGM
		GetNode<AudioStreamPlayer>("DeathSound").Play(); // 播放死亡音效

		GetNode<Timer>("MobTimer").Stop(); // 停止生成敌人
		GetNode<Timer>("ScoreTimer").Stop(); // 停止分数增加
		GetNode<Hud>("HUD").ShowGameOver(); // 让 UI 显示“游戏结束”信息
	}

	public void NewGame()
	{
		_score = 0;

		GetTree().CallGroup("mobs", Node.MethodName.QueueFree); // 移除所有 mob（前提是设置好了 group）

		var player = GetNode<Player>("Player");
		var startPosition = GetNode<Marker2D>("StartPosition");
		player.Start(startPosition.Position);

		GetNode<Timer>("StartTimer").Start();

		var hud = GetNode<Hud>("HUD");
		hud.UpdateScore(_score);
		hud.ShowMessage("Get Ready!"); // 让 UI 显示 GetReady 信息

		GetNode<AudioStreamPlayer>("Music").Play();
	}

	// 分数自增
	private void OnScoreTimerTimeout()
	{
		_score++;
		GetNode<Hud>("HUD").UpdateScore(_score); // 更新 UI 显示的分数
	}

	// “开始”定时器完成后做的操作
	private void OnStartTimerTimeout()
	{
		GetNode<Timer>("MobTimer").Start(); // 开始生成 Mob
		GetNode<Timer>("ScoreTimer").Start(); // 开始自增 Score
	}

	// mob 生成逻辑
	private void OnMobTimerTimeout()
	{
		// 创建 mob 场景的实例
		Mob mob = MobScene.Instantiate<Mob>();

		// 在 Path 上随机一个位置
		var mobSpawnLocation = GetNode<PathFollow2D>("MobPath/MobSpawnLocation");
		mobSpawnLocation.ProgressRatio = GD.Randf();

		// 右转 90 度（朝屏幕内）
		float direction = mobSpawnLocation.Rotation + Mathf.Pi / 2;

		// 设置 mob 的位置到该位置
		mob.Position = mobSpawnLocation.Position;

		// 为方向添加一些随机性
		direction += (float)GD.RandRange(-Mathf.Pi / 4, Mathf.Pi / 4);
		mob.Rotation = direction;

		// 随机速度
		var velocity = new Vector2((float)GD.RandRange(150.0, 250.0), 0);
		mob.LinearVelocity = velocity.Rotated(direction);

		// 添加为 main 的子节点（实际生成）
		AddChild(mob);
	}
}
