(defun Cos ()
  (setq a 3))

(defun double (number)
  (* 2 number))

(double 2)

(defun double-intercative (number)
  (interactive "p")
  (message "%d" (* 2 number)))

(defun greater-than-fill (number)
  (if (> fill-column number)
      (message "fill-column %d is greater than %d" fill-column number)
    (message "fill-column %d is lesser or equal than %d" fill-column number)))

(greater-than-fill 72)

(fill-column)

(defun simplified-beginning-of-buffer ()
  "Move point to the beginning of the buffer;
leave mark at previous position."
  (interactive)
  (push-mark)
  (goto-char (point-min)))
