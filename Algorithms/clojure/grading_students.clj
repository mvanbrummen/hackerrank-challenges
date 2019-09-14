

;
; Complete the 'gradingStudents' function below.
;
; The function is expected to return an INTEGER_ARRAY.
; The function accepts INTEGER_ARRAY grades as parameter.
;

(defn gradingStudents [grades]
    (map (fn [g]
           (let [diff (- 5 (mod g 5))]
             (if (and (< diff 3) (>= g 38)) (+ diff g) g))) grades))
  
  (def fptr (get (System/getenv) "OUTPUT_PATH"))
  
  (def grades-count (Integer/parseInt (clojure.string/trim (read-line))))
  
  (def grades [])
  
  (doseq [_ (range grades-count)]
      (def grades (conj grades (Integer/parseInt (clojure.string/trim (read-line)))))
  )
  
  (def result (gradingStudents grades))
  
  (spit fptr (clojure.string/join "\n" result) :append true)
  ;; (spit fptr "\n" :append true)
  