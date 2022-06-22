(defun multiply-by-seven (number)
  "Multiply NUMBER by seven."
  (* 7 number))

(multiply-by-seven 6)

(defun multiply-by-seven (number)
  "Multiply NUMBER by seven."
  (interactive "p")
  (message "The result is %d" (* 7 number)))

(let ((zebra 'stripes)
      (tiger 'fierce))
  (message "One kind of animal has %s and another is %s."
	   zebra tiger))

(let ((birch 3)
      pine
      fir
      (oak 'some))
  (message
   "Here are %d variables with %s, %s, and %s value."
   birch pine fir oak))

(if (> 5 4)
    (message "5 is greater than 4!"))

(defun type-of-animal (characteristic)
  "Print message in echo area depending on CHARACTERISTIC.
If the CHARACTERISTIC is the symbol 'fierce',
then warn of a tiger."
  (if (equal characteristic 'fierce)
      (message "It's a tiger!")))

(type-of-animal 'zebra)
(type-of-animal 'fierce)

(if (> 4 5)
    (message "4 falsely greater than 5!")
  (message "4 is not greater than 5!"))

(defun type-of-animal-2 (characteristic)
  "Print message in echo area depending on CHARACTERISTIC.
If the CHARACTERISTIC is the symbol 'fierce',
then warn of a tiger."
  (if (equal characteristic 'fierce)
      (message "It's a tiger!")
    (message "It's not a fierce")))

(type-of-animal-2 'zebra)
(type-of-animal-2 'fierce)

(if 4
    'true
  'false)

(if nil
    'true
  'false)

(> 5 4)

(> 4 5)

(let ((foo (buffer-name))
      (bar (buffer-size)))
  (message
   "This buffer is %s and has %d characters."
   foo bar))

(message "We are %d characters into this buffer."
	 (- (point)
	    (save-excursion
	     (goto-char (point-min)))))

(if (= 24 emacs-major-version)
    (message "This is version 24 Emacs")
  (message "This is not version 24 Emacs"))
