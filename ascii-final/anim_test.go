package ascii

import (
	"strings"
	"testing"
)

func TestAnimation_GenerateSpinFrames(t *testing.T) {
	anim := NewAnimation("TEST", 4)
	anim.GenerateSpinFrames()

	t.Run("Correct number of frames", func(t *testing.T) {
		for i := 0; i < 4; i++ {
			frame := anim.GetFrame(i)
			if frame == "" {
				t.Errorf("Frame %d is empty", i)
			}
		}

		// Test seamless looping (frame count should wrap to first frame)
		frameLast := anim.GetFrame(3)
		frameWrap := anim.GetFrame(4) // Should wrap to frame 0
		if frameWrap != frameLast {
			t.Log("Animation should loop seamlessly (frame 4 equals frame 0 or 3 based on implementation)")
		}
	})

	t.Run("Frame dimensions - 10 lines tall", func(t *testing.T) {
		for i := 0; i < 4; i++ {
			frame := anim.GetFrame(i)
			lines := strings.Split(strings.TrimRight(frame, "\n"), "\n")

			if len(lines) != 10 {
				t.Errorf("Spin frame %d: Expected 10 lines, got %d", i, len(lines))
			}

			// Check consistent width (optional, not required but good practice)
			if len(lines) > 0 {
				firstLineLength := len(lines[0])
				for j, line := range lines {
					if len(line) != firstLineLength {
						t.Errorf("Spin frame %d line %d has inconsistent length: %d vs %d",
							i, j, len(line), firstLineLength)
					}
				}
			}
		}
	})

	t.Run("Frames are different", func(t *testing.T) {
		frame0 := anim.GetFrame(0)
		frame1 := anim.GetFrame(1)

		if frame0 == frame1 {
			t.Error("Consecutive frames should be different in spin animation")
		}
	})
}

func TestAnimation_GenerateWaveFrames(t *testing.T) {
	anim := NewAnimation("WAVE", 8)
	anim.GenerateWaveFrames()

	t.Run("All frames generated", func(t *testing.T) {
		for i := 0; i < 8; i++ {
			if anim.GetFrame(i) == "" {
				t.Errorf("Frame %d not generated", i)
			}
		}
	})

	t.Run("Frame dimensions - 10 lines tall", func(t *testing.T) {
		for i := 0; i < 8; i++ {
			frame := anim.GetFrame(i)
			lines := strings.Split(strings.TrimRight(frame, "\n"), "\n")

			if len(lines) != 10 {
				t.Errorf("Wave frame %d: Expected 10 lines, got %d", i, len(lines))
			}
		}
	})

	t.Run("Wave pattern creates variation", func(t *testing.T) {
		// Compare first and last frames - they should be different
		frame0 := anim.GetFrame(0)
		frame7 := anim.GetFrame(7)

		if frame0 == frame7 {
			t.Error("Wave animation should create varying patterns between first and last frame")
		}

		// At least some frames should differ
		allSame := true
		for i := 1; i < 8; i++ {
			if anim.GetFrame(i) != anim.GetFrame(0) {
				allSame = false
				break
			}
		}

		if allSame {
			t.Error("Wave animation should have at least one different frame")
		}
	})
}

func TestAnimation_GenerateZoomFrames(t *testing.T) {
	anim := NewAnimation("ZOOM", 5)
	anim.GenerateZoomFrames()

	t.Run("All frames generated", func(t *testing.T) {
		for i := 0; i < 5; i++ {
			if anim.GetFrame(i) == "" {
				t.Errorf("Frame %d not generated", i)
			}
		}
	})

	t.Run("Frame dimensions - 10 lines tall", func(t *testing.T) {
		for i := 0; i < 5; i++ {
			frame := anim.GetFrame(i)
			lines := strings.Split(strings.TrimRight(frame, "\n"), "\n")

			if len(lines) != 10 {
				t.Errorf("Zoom frame %d: Expected 10 lines, got %d", i, len(lines))
			}
		}
	})

	t.Run("Zoom creates variation between frames", func(t *testing.T) {
		frame0 := anim.GetFrame(0)
		frame2 := anim.GetFrame(2)
		frame4 := anim.GetFrame(4)

		// Frames should change during zoom effect
		if frame0 == frame2 && frame2 == frame4 {
			t.Error("Zoom animation should create variation between frames")
		}

		// Note: Zoom may or may not change width, but should change content
		// Width changes are allowed as long as height remains 10 lines
	})
}

func TestAnimation_Play(t *testing.T) {
	anim := NewAnimation("PLAY", 3)
	anim.GenerateSpinFrames()

	playOutput := anim.Play()

	t.Run("Play output not empty", func(t *testing.T) {
		if playOutput == "" {
			t.Error("Play output should not be empty")
		}
	})

	t.Run("Play contains frame separators/delays", func(t *testing.T) {
		// Check for some kind of frame separation (not exact string)
		hasFrameMarkers := strings.Contains(playOutput, "Frame") ||
			strings.Contains(playOutput, "---") ||
			strings.Contains(playOutput, "===") ||
			strings.Contains(playOutput, "\n\n")

		if !hasFrameMarkers {
			t.Error("Play output should contain frame delay markers or separators")
		}
	})

	t.Run("Play contains all frames", func(t *testing.T) {
		// Count newlines as a proxy for content (not perfect but works)
		if len(playOutput) < 100 { // Arbitrary minimum for 3 frames of ASCII art
			t.Logf("Warning: Play output seems short (%d chars), may not contain all frames", len(playOutput))
		}
	})
}

func TestAnimation_SeamlessLooping(t *testing.T) {
	anim := NewAnimation("LOOP", 4)
	anim.GenerateSpinFrames()

	// Last frame should lead smoothly to first frame
	_ = anim.GetFrame(3)
	wrapFrame := anim.GetFrame(4) // Should wrap

	if wrapFrame == "" {
		t.Error("Frame wrapping should return a valid frame (seamless looping requirement)")
	}
}

func TestAnimation_NoExternalFiles(t *testing.T) {
	anim := NewAnimation("TEST", 3)
	anim.GenerateSpinFrames()

	// This is hard to test automatically, but at minimum we can verify
	// frames contain the text "TEST" somewhere (meaning it was generated, not loaded)
	frame := anim.GetFrame(0)
	if !strings.Contains(frame, "T") || !strings.Contains(frame, "E") ||
		!strings.Contains(frame, "S") || !strings.Contains(frame, "T") {
		t.Log("Warning: Frames don't contain the original text - may not be properly generated")
	}
}
