using Godot;

public partial class Hud : CanvasLayer
{
	[Signal]
	public delegate void StartGameEventHandler(); // 本 scene 作为节点时会 emit 的信号

	public void ShowMessage(string text)
	{
		var message = GetNode<Label>("Message");
		message.Text = text;
		message.Show();

		// MessageTimer 是编辑器中定义的 one shot 2s 定时器，作用是使每个信息显示 2s
		GetNode<Timer>("MessageTimer").Start();
	}
	
	async public void ShowGameOver()
	{
		ShowMessage("Game Over");

		var messageTimer = GetNode<Timer>("MessageTimer");
		await ToSignal(messageTimer, Timer.SignalName.Timeout); // 阻塞式等待“信息显示 2s”的时间过去

		var message = GetNode<Label>("Message");
		message.Text = "Dodge the Creeps!"; // 恢复为游戏标题
		message.Show();

		// 创建临时的一次性定时器（结束后会自动释放内存），并阻塞式等待它
		await ToSignal(GetTree().CreateTimer(1.0), SceneTreeTimer.SignalName.Timeout);
		GetNode<Button>("StartButton").Show(); // 显示开始按钮
	}

	// 更新分数显示
	public void UpdateScore(int score)
	{
		GetNode<Label>("ScoreLabel").Text = score.ToString();
	}

	// 通过信号编辑器连接，让 StartButton 的按下能触发本函数
	private void OnStartButtonPressed()
	{
		GetNode<Button>("StartButton").Hide(); // 让开始按钮消失
		EmitSignal(SignalName.StartGame); // 触发本节点的“开始游戏”信号，让 main 场景知道
	}

	// 通过信号编辑器连接，让 MessageTimer 时间到后触发本函数
	private void OnMessageTimerTimeout()
	{
		GetNode<Label>("Message").Hide(); // 隐藏 Message 节点
	}
}
