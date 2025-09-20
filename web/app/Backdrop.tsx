// 画面全体を覆う黒色透明オーバーレイ
export function Backdrop({ zIndex }: { zIndex: number }) {
  return (
    <div style={{
      position: 'fixed',
      top: 0,
      left: 0,
      width: '100vw',
      height: '100vh',
      backgroundColor: 'rgba(0,0,0,0.5)',
      zIndex: zIndex,
      pointerEvents: 'auto',
    }} />
  );
}
