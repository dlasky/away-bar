package main

// func InitBattery() (gtk.IWidget, error) {
// 	batLabel, err := gtk.LabelNew("")
// 	if err != nil {
// 		return nil, err
// 	}

// 	go func() {
// 		for {
// 			batData, err := battery.Get(0)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			s := fmt.Sprintf("bat: %.0f %%", batData.Current/batData.Full*100)
// 			sh = glib.IdleAdd(batLabel.SetText, s)
// 			sh.Add()
// 			// if err != nil {
// 			// 	log.Fatal("IdleAdd() failed:", err)
// 			// }
// 			time.Sleep(60 * time.Second)
// 		}
// 	}()

// 	return batLabel, nil
// }
