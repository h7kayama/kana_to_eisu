package main

import (
	"fmt"
	"os/exec"
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
			cmd := exec.Command("osascript", "-e", `tell application "System Events" to key code {102}`)
			if err := cmd.Run(); err != nil {
				fmt.Printf("Failed to execute key code: %v\n", err)
			}
			lastInputTime = C.int64_t(currentTime)
		}
	}
}
