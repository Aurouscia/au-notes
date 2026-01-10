# 细微的位置差异
我发现，同一个元素，有 transform: rotate(1deg) 和没有时，位置有细微的差别（不到1px）

## 解释
这是浏览器在渲染时做「像素对齐（pixel snapping）」导致的。
带 transform:rotate(…) 的元素会被提升为独立的合成层（composited layer），光栅化时浏览器会把图层内容对齐到整像素，防止子像素模糊；而不加 transform 的普通盒子仍然走普通的布局/绘制流程，两者在舍入规则上略有差异，于是你会看到不到 1 px 的偏移。

## 解决方法
让元素**始终**走合成层，这样无论有没有旋转，舍入规则都一致，以下两个属性二选一确保常驻生效即可：
- will-change: transform;
- rotate(0.0001deg);