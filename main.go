package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
	"unsafe"
)

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa -framework ApplicationServices
#import <ApplicationServices/ApplicationServices.h>
#import <Cocoa/Cocoa.h>

static CGEventRef eventCallback(CGEventTapProxy proxy, CGEventType type, CGEventRef event, void *refcon) {
    if (type == kCGEventKeyDown || type == kCGEventFlagsChanged) {
        int64_t *lastInput = (int64_t*)refcon;
        *lastInput = time(NULL);
    }
    return event;
}

static void* startEventTap(void *refcon) {
    CFRunLoopSourceRef runLoopSource;
    CGEventMask eventMask = CGEventMaskBit(kCGEventKeyDown) | CGEventMaskBit(kCGEventFlagsChanged);
    CFMachPortRef eventTap = CGEventTapCreate(
        kCGSessionEventTap,
        kCGHeadInsertEventTap,
        kCGEventTapOptionDefault,
        eventMask,
        eventCallback,
        refcon
    );

    if (!eventTap) {
        printf("Failed to create event tap\n");
        return NULL;
    }

    runLoopSource = CFMachPortCreateRunLoopSource(kCFAllocatorDefault, eventTap, 0);
    CFRunLoopAddSource(CFRunLoopGetCurrent(), runLoopSource, kCFRunLoopCommonModes);
    CGEventTapEnable(eventTap, true);
    CFRunLoopRun();
    return NULL;
}
*/
import "C"

func main() {
	fmt.Println("Started monitoring keyboard input. Press Ctrl+C to exit.")

	var lastInputTime C.int64_t
	lastInputTime = C.int64_t(time.Now().Unix())

	// Start event tap
	go C.startEventTap(unsafe.Pointer(&lastInputTime))

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		currentTime := time.Now().Unix()
		if currentTime-int64(lastInputTime) >= 3 {
			// Get current IME mode
			cmd := exec.Command("im-select")
			currentIMEMode, err := cmd.Output()
			if err != nil {
				fmt.Printf("Failed to get current IME mode: %v\n", err)
				continue
			}

			// Switch to English input mode only if current mode is Japanese
			if strings.TrimSpace(string(currentIMEMode)) == "com.apple.inputmethod.Kotoeri.RomajiTyping.Japanese" {
				cmd = exec.Command("im-select", "com.apple.keylayout.ABC")
				if err := cmd.Run(); err != nil {
					fmt.Printf("Failed to switch IME: %v\n", err)
				}
			}
			lastInputTime = C.int64_t(currentTime)
		}
	}
}
