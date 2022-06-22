(message "I co %s" "dziala?")
(message "I co %s" 3)
`(message "I co %s" dziala?)
(message "The name of this buffer is: %s." (buffer-name))
(message "I co %s" 3.1)
(message "I co %d" 3.1)
(set `a 4)
(set 'b 3)
a
(message "The value of fill-column is %d." fill-column)
(message "There is %d %s in the office!"
	 (- fill-column 14) "pink elefants")
(message "He saw %d %s"
	 (- fill-column 32)
	 (concat "red "
		 (substring "The quick brown foxes jumped." 16 21)
		 " leaping."))

(setq counter 0)

(setq counter (+ counter 2))

counter

(message "Proste.")