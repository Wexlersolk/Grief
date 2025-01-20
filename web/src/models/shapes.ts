// shapes.ts

export class Rectangle {
  constructor(public x: number, public y: number, public width: number = 0, public height: number = 0) {}

  update(x: number, y: number) {
    this.width = x - this.x;
    this.height = y - this.y;
  }

  draw(ctx: CanvasRenderingContext2D) {
    ctx.rect(this.x, this.y, this.width, this.height);
    ctx.stroke();
  }
}

export class Arrow {
  constructor(public x: number, public y: number, public endX: number = x, public endY: number = y) {}

  update(x: number, y: number) {
    this.endX = x;
    this.endY = y;
  }

  draw(ctx: CanvasRenderingContext2D) {
    ctx.moveTo(this.x, this.y);
    ctx.lineTo(this.endX, this.endY);
    ctx.stroke();
    ctx.beginPath();
    ctx.moveTo(this.endX, this.endY);
    ctx.lineTo(this.endX - 10, this.endY - 10); // Arrow head
    ctx.moveTo(this.endX, this.endY);
    ctx.lineTo(this.endX + 10, this.endY - 10); // Arrow head
    ctx.stroke();
  }
}

export class Line {
  constructor(public x: number, public y: number) {}

  update(x: number, y: number) {
    this.x = x;
    this.y = y;
  }

  draw(ctx: CanvasRenderingContext2D) {
    ctx.lineTo(this.x, this.y);
    ctx.stroke();
  }
}

