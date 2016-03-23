package main

import (
	"fmt"

	"github.com/PieterD/glimmer/gli"
	"github.com/PieterD/glimmer/mat"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

const d2r = 3.14159 * 2 / 360

type Crane struct {
	posBase        mgl32.Vec3
	angBase        float32
	posBaseLeft    mgl32.Vec3
	posBaseRight   mgl32.Vec3
	scaleBaseZ     float32
	angUpperArm    float32
	sizeUpperArm   float32
	posLowerArm    mgl32.Vec3
	angLowerArm    float32
	lenLowerArm    float32
	widthLowerArm  float32
	posWrist       mgl32.Vec3
	angWristRoll   float32
	angWristPitch  float32
	lenWrist       float32
	widthWrist     float32
	posLeftFinger  mgl32.Vec3
	posRightFinger mgl32.Vec3
	angFingerOpen  float32
	lenFinger      float32
	widthFinger    float32
	angLowerFinger float32
	p              *Profile
}

func NewCrane(p *Profile) *Crane {
	return &Crane{
		posBase:        mgl32.Vec3{3, -5, -40},
		angBase:        -45,
		posBaseLeft:    mgl32.Vec3{2, 0, 0},
		posBaseRight:   mgl32.Vec3{-2, 0, 0},
		scaleBaseZ:     3,
		angUpperArm:    -78,
		sizeUpperArm:   9,
		posLowerArm:    mgl32.Vec3{0, 0, 8},
		angLowerArm:    101,
		lenLowerArm:    5,
		widthLowerArm:  1.5,
		posWrist:       mgl32.Vec3{0, 0, 5},
		angWristRoll:   0,
		angWristPitch:  67.5,
		lenWrist:       2,
		widthWrist:     2,
		posLeftFinger:  mgl32.Vec3{1, 0, 1},
		posRightFinger: mgl32.Vec3{-1, 0, 1},
		angFingerOpen:  27,
		lenFinger:      2,
		widthFinger:    0.5,
		angLowerFinger: 45,
		p:              p,
	}
}

func (crane *Crane) Draw() {
	stack := mat.NewStack()
	stack.Deg()
	stack.TranslateV(crane.posBase)
	stack.RotateY(crane.angBase)
	stack.Multiply()

	stack.Safe(crane.base)
	stack.Safe(crane.arm)
}

func (crane *Crane) base(stack *mat.Stack) {
	stack.Safe(func(stack *mat.Stack) {
		stack.TranslateV(crane.posBaseLeft).Multiply()
		stack.Scale(1, 1, crane.scaleBaseZ).Multiply()
		crane.put(stack)
	})
	stack.Safe(func(stack *mat.Stack) {
		stack.TranslateV(crane.posBaseRight).Multiply()
		stack.Scale(1, 1, crane.scaleBaseZ).Multiply()
		crane.put(stack)
	})
}

func (crane *Crane) arm(stack *mat.Stack) {
	stack.RightMul(mgl32.HomogRotate3DX(crane.angUpperArm * d2r))
	stack.Safe(func(stack *mat.Stack) {
		stack.RightMul(mgl32.Translate3D(0, 0, crane.sizeUpperArm/2-1))
		stack.RightMul(mgl32.Scale3D(1, 1, crane.sizeUpperArm/2))
		crane.put(stack)
	})
	stack.Safe(crane.lowerarm)
}

func (crane *Crane) lowerarm(stack *mat.Stack) {
	stack.RightMul(translate(crane.posLowerArm))
	stack.RightMul(mgl32.HomogRotate3DX(crane.angLowerArm * d2r))
	stack.Safe(func(stack *mat.Stack) {
		stack.RightMul(mgl32.Translate3D(0, 0, crane.lenLowerArm/2))
		stack.RightMul(mgl32.Scale3D(crane.widthLowerArm/2, crane.widthLowerArm/2, crane.lenLowerArm/2))
		crane.put(stack)
	})
	stack.Safe(crane.wrist)
}

func (crane *Crane) wrist(stack *mat.Stack) {
	stack.RightMul(translate(crane.posWrist))
	stack.RightMul(mgl32.HomogRotate3DZ(crane.angWristRoll * d2r))
	stack.RightMul(mgl32.HomogRotate3DX(crane.angWristPitch * d2r))
	stack.Safe(func(stack *mat.Stack) {
		stack.RightMul(mgl32.Scale3D(crane.widthWrist/2, crane.widthWrist/2, crane.lenWrist/2))
		crane.put(stack)
	})
	stack.Safe(crane.leftfinger)
	stack.Safe(crane.rightfinger)
}

func (crane *Crane) leftfinger(stack *mat.Stack) {
	stack.RightMul(translate(crane.posLeftFinger))
	stack.RightMul(mgl32.HomogRotate3DY(crane.angFingerOpen * d2r))
	stack.Safe(func(stack *mat.Stack) {
		stack.RightMul(mgl32.Translate3D(0, 0, crane.lenFinger/2))
		stack.RightMul(mgl32.Scale3D(crane.widthFinger/2, crane.widthFinger/2, crane.lenFinger/2))
		crane.put(stack)
	})
	stack.Safe(func(stack *mat.Stack) {
		stack.RightMul(mgl32.Translate3D(0, 0, crane.lenFinger))
		stack.RightMul(mgl32.HomogRotate3DY(-crane.angLowerFinger * d2r))
		stack.Safe(func(stack *mat.Stack) {
			stack.RightMul(mgl32.Translate3D(0, 0, crane.lenFinger/2))
			stack.RightMul(mgl32.Scale3D(crane.widthFinger/2, crane.widthFinger/2, crane.lenFinger/2))
			crane.put(stack)
		})
	})
}

func (crane *Crane) rightfinger(stack *mat.Stack) {
	stack.RightMul(translate(crane.posRightFinger))
	stack.RightMul(mgl32.HomogRotate3DY(-crane.angFingerOpen * d2r))
	stack.Safe(func(stack *mat.Stack) {
		stack.RightMul(mgl32.Translate3D(0, 0, crane.lenFinger/2))
		stack.RightMul(mgl32.Scale3D(crane.widthFinger/2, crane.widthFinger/2, crane.lenFinger/2))
		crane.put(stack)
	})
	stack.Safe(func(stack *mat.Stack) {
		stack.RightMul(mgl32.Translate3D(0, 0, crane.lenFinger))
		stack.RightMul(mgl32.HomogRotate3DY(crane.angLowerFinger * d2r))
		stack.Safe(func(stack *mat.Stack) {
			stack.RightMul(mgl32.Translate3D(0, 0, crane.lenFinger/2))
			stack.RightMul(mgl32.Scale3D(crane.widthFinger/2, crane.widthFinger/2, crane.lenFinger/2))
			crane.put(stack)
		})
	})
}

func (crane *Crane) adjustBase(dir float32) {
	crane.angBase += 11.25 * dir
}

func (crane *Crane) adjustUpperArm(dir float32) {
	crane.angUpperArm += 11.25 * dir
}

func (crane *Crane) adjustLowerArm(dir float32) {
	crane.angLowerArm += 11.25 * dir
}

func (crane *Crane) adjustWristPitch(dir float32) {
	crane.angWristPitch += 11.25 * dir
}

func (crane *Crane) adjustWristRoll(dir float32) {
	crane.angWristRoll += 11.25 * dir
}

func (crane *Crane) adjustFingerOpen(dir float32) {
	crane.angFingerOpen += 9 * dir
}

func (crane *Crane) put(stack *mat.Stack) {
	p := crane.p
	m := stack.Peek()
	p.modelToCameraMatrix.Float(m[:]...)
	gli.Draw(p.program, p.vao, rectObject)
}

func translate(v mgl32.Vec3) mgl32.Mat4 {
	x, y, z := v.Elem()
	return mgl32.Translate3D(x, y, z)
}

func (p *Profile) EventRune(w *glfw.Window, char rune) {
	switch char {
	case 'a':
		p.crane.adjustBase(1)
	case 'd':
		p.crane.adjustBase(-1)
	case 'w':
		p.crane.adjustUpperArm(-1)
	case 's':
		p.crane.adjustUpperArm(1)
	case 'r':
		p.crane.adjustLowerArm(-1)
	case 'f':
		p.crane.adjustLowerArm(1)
	case 't':
		p.crane.adjustWristPitch(-1)
	case 'g':
		p.crane.adjustWristPitch(1)
	case 'z':
		p.crane.adjustWristRoll(1)
	case 'c':
		p.crane.adjustWristRoll(-1)
	case 'q':
		p.crane.adjustFingerOpen(1)
	case 'e':
		p.crane.adjustFingerOpen(-1)
	}
	fmt.Printf("base %f\nupper %f\nlower %f\nwpitch %f\nwroll %f\nfinger %f\n\n", p.crane.angBase, p.crane.angUpperArm, p.crane.angLowerArm, p.crane.angWristPitch, p.crane.angWristRoll, p.crane.angFingerOpen)
}
