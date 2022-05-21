package main

import (
	"fmt"
	"github.com/inancgumus/screen"
	"sort"
	"strconv"
)

func menu_loop() {
	const ITEM_SHOWINFO = 0
	const ITEM_PATCH = 1
	const ITEM_PATCH_PROCS = 2
	const ITEM_EXIT = 3

	items := map[byte]string{
		ITEM_SHOWINFO:    "Show info",
		ITEM_PATCH:       "Patch (default flow)",
		ITEM_PATCH_PROCS: "Smart (but not clever) patch everything that is running",
		ITEM_EXIT:        "Exit",
	}

	var itemKeysIndex []byte
	for i, _ := range items {
		itemKeysIndex = append(itemKeysIndex, i)
	}

	sort.SliceStable(itemKeysIndex, func(i, j int) bool {
		return i < j
	})

	func() {
		var selected_item int
		for {
			screen.Clear()
			screen.MoveTopLeft()

			for _, i := range itemKeysIndex {
				fmt.Printf("%d. %s\n", i, items[i])
			}

			var inbuf []byte
			inbuf, _, _ = stdin.ReadLine()
			selected_item, _ = strconv.Atoi(string(inbuf))

			if selected_item == ITEM_EXIT {
				break
			}

			switch selected_item {
			case ITEM_SHOWINFO:
				item_show_info()
				break
			case ITEM_PATCH:
				item_patch()
				break
			case ITEM_PATCH_PROCS:
				item_patch_procs()
				break
			default:
				fmt.Println("Incorrect choice")
				break
			}

			delay()
		}
	}()
}
