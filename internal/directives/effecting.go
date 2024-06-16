package directives

// type datadirectives struct {
// 	named      []models.WildcardKeypath
// 	parent     []models.WildcardKeypath
// 	typeassign []models.WildcardKeypath
// }

// type typedirectives struct {

// }

// type effectCollection map[models.FlattenKeypath]keypathdirs

// func (ewckp keypathdirs) Print() {
// 	if len(ewckp.named) > 0 {
// 		fmt.Printf("    named: %v\n", ewckp.named)
// 	}
// 	if len(ewckp.parent) > 0 {
// 		fmt.Printf("    parent: %v\n", ewckp.parent)
// 	}
// 	if len(ewckp.typeassign) > 0 {
// 		fmt.Printf("    type: %v\n", ewckp.typeassign)
// 	}
// }

// func caseInsensitiveCompareKeypaths(a, b models.FlattenKeypath) int {
// 	if strings.ToLower(string(a)) < strings.ToLower(string(b)) {
// 		return -1
// 	} else {
// 		return +1
// 	}
// }

// func (ec effectCollection) Print() {
// 	fmt.Println("rules grouped in effecting keypaths:")
// 	sorted := maps.Keys(ec)
// 	slices.SortFunc(sorted, caseInsensitiveCompareKeypaths)
// 	for _, kp := range sorted {
// 		dirs := ec[kp]
// 		fmt.Printf("  %s\n", kp)
// 		dirs.Print()
// 	}
// }

// func effecting() {

// 	directives := effectCollection{}
// 	for kp, wckps := range effects {
// 		kpdirs := keypathdirs{}
// 		for _, wckp := range wckps {
// 			wckpdirs := (*b.Df)[wckp]
// 			if wckpdirs.Named != "" {
// 				kpdirs.named = append(kpdirs.named, wckp)
// 			}
// 			if wckpdirs.Parent != "" {
// 				kpdirs.parent = append(kpdirs.parent, wckp)
// 			}
// 			if wckpdirs.Type != "" {
// 				kpdirs.typeassign = append(kpdirs.typeassign, wckp)
// 			}
// 			directives[kp] = kpdirs
// 		}
// 	}

// }
