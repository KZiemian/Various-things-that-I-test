(defun mark-whole-buffer-1 ()
  "Put point at beginning and mark at end of buffer
You probably shlould not use this function in Lisp programs;
it is usually a mistake for a Lisp function tu use any subrutine
that uses or sets the mark."
  (interactive)
  (push-mark (point))
  (push-mark (point-max) nil t)
  (goto-char (point-min)))

(push-mark (point-max))

(mark-whole-buffer-1)

(defun append-to-buffer-1 (buffer start end)
  "Append to specified buffer the text of the region.
It is inserted into that buffer  before its point.

When calling from a program, give three arguments:
BUFFER (or buffer name), START and END.
START and END specify the portion of the current buffer to be copied."
  (interactive
   (list (read-buffer "Append to buffer: "
		      (other-buffer (current-buffer) t))
	 (region-beginning) (region-end)))
  (let ((oldbuf (current-buffer)))
    (save-excursion
      (let* ((append-to (get-buffer-create buffer))
	     (windows (get-buffer-window-list append-to t t))
	     point)
	(set-buffer append-to)
	(setq point (point))
	(barf-it-buffer-read-only)
	(insert-buffer-substring oldbuf start end)
	(dolist (window windows)
	  (when (= (window-point window) point)
	    (set-window-point window (point))))))))
