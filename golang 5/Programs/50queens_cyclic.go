package main

import "fmt"

const N = 50

type queen struct {
	          col int
	          row int
	       }     

func abs(x int) int {
	if x >= 0  {
		return x 
	} else {
		return -x
	}
}			

func Connected (q1, q2 queen) bool {
	return q1.col == q2.col ||
	       q1.row == q2.row ||
	       abs(q1.col - q2.col) == abs (q1.row - q2.row) 
}

func Conflict (qs []queen, q queen) bool {
	for _, q2 := range qs {
		if Connected (q, q2) { return true }	
	}	
	return false
}	

func main() {
	var Queens [N]queen
	for col0:= 0; col0 < N; col0++ {
		Queens[0] = queen{col0, 0}///--------------------------1
		if Conflict(Queens[:0], Queens[0]) {continue}
		for col1:= 0; col1 < N; col1++ {
			Queens[1] = queen{col1, 1}///--------------------------2
			if Conflict(Queens[:1], Queens[1]) {continue}
			for col2:= 0; col2 < N; col2++ {
				Queens[2] = queen{col2, 2}///--------------------------3
				if Conflict(Queens[:2], Queens[2]) {continue}
				for col3:= 0; col3 < N; col3++ {
					Queens[3] = queen{col3, 3}///--------------------------4
					if Conflict(Queens[:3], Queens[3]) {continue}
					for col4:= 0; col4 < N; col4++ {
						Queens[4] = queen{col4, 4}///--------------------------5
						if Conflict(Queens[:4], Queens[4]) {continue}
						for col5:= 0; col5 < N; col5++ {
							Queens[5] = queen{col5, 5}///--------------------------6
							if Conflict(Queens[:5], Queens[5]) {continue}
							for col6:= 0; col6 < N; col6++ {
								Queens[6] = queen{col6, 6}///--------------------------7
								if Conflict(Queens[:6], Queens[6]) {continue}
								for col7:= 0; col7 < N; col7++ {
									Queens[7] = queen{col7, 7}///--------------------------8
									if Conflict(Queens[:7], Queens[7]) {continue}
									for col8:= 0; col8 < N; col8++ {
										Queens[8] = queen{col8, 8}///--------------------------9
										if Conflict(Queens[:8], Queens[8]) {continue}
										for col9:= 0; col9 < N; col9++ {
											Queens[9] = queen{col9, 9}///--------------------------10
											if Conflict(Queens[:9], Queens[9]) {continue}
											for col10:= 0; col10 < N; col10++ {
												Queens[10] = queen{col10, 10}///--------------------------11
												if Conflict(Queens[:10], Queens[10]) {continue}
												for col11:= 0; col11 < N; col11++ {
													Queens[11] = queen{col11, 11}///--------------------------12
													if Conflict(Queens[:11], Queens[11]) {continue}
													for col12:= 0; col12 < N; col12++ {
														Queens[12] = queen{col12, 12}///--------------------------13
														if Conflict(Queens[:12], Queens[12]) {continue}
														for col13:= 0; col13 < N; col13++ {
															Queens[13] = queen{col13, 13}///--------------------------14
															if Conflict(Queens[:13], Queens[13]) {continue}
															for col14:= 0; col14 < N; col14++ {
																Queens[14] = queen{col14, 14}///--------------------------15
																if Conflict(Queens[:14], Queens[14]) {continue}
																for col15:= 0; col15 < N; col15++ {
																	Queens[15] = queen{col15, 15}///--------------------------16
																	if Conflict(Queens[:15], Queens[15]) {continue}
																	for col16:= 0; col16 < N; col16++ {
																		Queens[16] = queen{col16, 16}///--------------------------17
																		if Conflict(Queens[:16], Queens[16]) {continue}
																		for col17:= 0; col17 < N; col17++ {
																			Queens[17] = queen{col17, 17}///--------------------------18
																			if Conflict(Queens[:17], Queens[17]) {continue}
																			for col18:= 0; col18 < N; col18++ {
																				Queens[18] = queen{col18, 18}///--------------------------19
																				if Conflict(Queens[:18], Queens[18]) {continue}
																				for col19:= 0; col19 < N; col19++ {
																					Queens[19] = queen{col19, 19}///--------------------------20
																					if Conflict(Queens[:19], Queens[19]) {continue}
																					for col20:= 0; col20 < N; col20++ {
																						Queens[20] = queen{col20, 20}///--------------------------21
																						if Conflict(Queens[:20], Queens[20]) {continue}
																						for col21:= 0; col21 < N; col21++ {
																							Queens[21] = queen{col21, 21}///--------------------------22
																							if Conflict(Queens[:21], Queens[21]) {continue}
																							for col22:= 0; col22 < N; col22++ {
																								Queens[22] = queen{col22, 22}///--------------------------23
																								if Conflict(Queens[:22], Queens[22]) {continue}
																								for col23:= 0; col23 < N; col23++ {
																									Queens[23] = queen{col23, 23}///--------------------------24
																									if Conflict(Queens[:23], Queens[23]) {continue}
																									for col24:= 0; col24 < N; col24++ {
																										Queens[24] = queen{col24, 24}///--------------------------25
																										if Conflict(Queens[:24], Queens[24]) {continue}
																										for col25:= 0; col25 < N; col25++ {
																											Queens[25] = queen{col25, 25}///--------------------------26
																											if Conflict(Queens[:25], Queens[25]) {continue}
																											for col26:= 0; col26 < N; col26++ {
																												Queens[26] = queen{col26, 26}///--------------------------27
																												if Conflict(Queens[:26], Queens[26]) {continue}
																												for col27:= 0; col27 < N; col27++ {
																													Queens[27] = queen{col27, 27}///--------------------------28
																													if Conflict(Queens[:27], Queens[27]) {continue}
																													for col28:= 0; col28 < N; col28++ {
																														Queens[28] = queen{col28, 28}///--------------------------29
																														if Conflict(Queens[:28], Queens[28]) {continue}
																														for col29:= 0; col29 < N; col29++ {
																															Queens[29] = queen{col29, 29}///--------------------------30
																															if Conflict(Queens[:29], Queens[29]) {continue}
																															for col30:= 0; col30 < N; col30++ {
																																Queens[30] = queen{col30, 30}///--------------------------31
																																if Conflict(Queens[:30], Queens[30]) {continue}
																																for col31:= 0; col31 < N; col31++ {
																																	Queens[31] = queen{col31, 31}///--------------------------32
																																	if Conflict(Queens[:31], Queens[31]) {continue}
																																	for col32:= 0; col32 < N; col32++ {
																																		Queens[32] = queen{col32, 32}///--------------------------33
																																		if Conflict(Queens[:32], Queens[32]) {continue}
																																		for col33:= 0; col33 < N; col33++ {
																																			Queens[33] = queen{col33, 33}///--------------------------34
																																			if Conflict(Queens[:33], Queens[33]) {continue}
																																			for col34:= 0; col34 < N; col34++ {
																																				Queens[34] = queen{col34, 34}///--------------------------35
																																				if Conflict(Queens[:34], Queens[34]) {continue}
																																				for col35:= 0; col35 < N; col35++ {
																																					Queens[35] = queen{col35, 35}///--------------------------36
																																					if Conflict(Queens[:35], Queens[35]) {continue}
																																					for col36:= 0; col36 < N; col36++ {
																																						Queens[36] = queen{col36, 36}///--------------------------37
																																						if Conflict(Queens[:36], Queens[36]) {continue}
																																						for col37:= 0; col37 < N; col37++ {
																																							Queens[37] = queen{col37, 37}///--------------------------38
																																							if Conflict(Queens[:37], Queens[37]) {continue}
																																							for col38:= 0; col38 < N; col38++ {
																																								Queens[38] = queen{col38, 38}///--------------------------39
																																								if Conflict(Queens[:38], Queens[38]) {continue}
																																								for col39:= 0; col39 < N; col39++ {
																																									Queens[39] = queen{col39, 39}///--------------------------39
																																									if Conflict(Queens[:39], Queens[39]) {continue}
																																									for col40:= 0; col40 < N; col40++ {
																																										Queens[40] = queen{col40, 40}///--------------------------41
																																										if Conflict(Queens[:40], Queens[40]) {continue}
																																										for col41:= 0; col41 < N; col41++ {
																																											Queens[41] = queen{col41, 41}///--------------------------42
																																											if Conflict(Queens[:41], Queens[41]) {continue}
																																											for col42:= 0; col42 < N; col42++ {
																																												Queens[42] = queen{col42, 42}///--------------------------43
																																												if Conflict(Queens[:42], Queens[42]) {continue}
																																												for col43:= 0; col43 < N; col43++ {
																																													Queens[43] = queen{col43, 43}///--------------------------44
																																													if Conflict(Queens[:43], Queens[43]) {continue}
																																													for col44:= 0; col44 < N; col44++ {
																																														Queens[44] = queen{col44, 44}///--------------------------45
																																														if Conflict(Queens[:44], Queens[44]) {continue}
																																														for col45:= 0; col45 < N; col45++ {
																																															Queens[45] = queen{col45, 45}///--------------------------46
																																															if Conflict(Queens[:45], Queens[45]) {continue}
																																															for col46:= 0; col46 < N; col46++ {
																																																Queens[46] = queen{col46, 46}///--------------------------47
																																																if Conflict(Queens[:46], Queens[46]) {continue}
																																																for col47:= 0; col47 < N; col47++ {
																																																	Queens[47] = queen{col47, 47}///--------------------------48
																																																	if Conflict(Queens[:47], Queens[47]) {continue}
																																																	for col48:= 0; col48 < N; col48++ {
																																																		Queens[48] = queen{col48, 48}///--------------------------49
																																																		if Conflict(Queens[:48], Queens[48]) {continue}
																																																		for col49:= 0; col49 < N; col49++ {
																																																			Queens[49] = queen{col49, 49}///--------------------------50
																																																			if Conflict(Queens[:49], Queens[49]) {continue}
																																																			for _, q := range Queens {
																																																				fmt.Printf("%c%d ", q.col+'a', q.row + 1)
																																																			}
																																																			fmt.Println()
																																																		}
																																																	}
																																																}
																																															}
																																														}
																																													}
																																												}
																																											}
																																										}
																																									}
																																								}
																																							}
																																						}
																																					}
																																				}
																																			}
																																		}
																																	}
																																}
																															}
																														}
																													}
																												}
																											}
																										}
																									}
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
