

; Complete the caesarCipher function below.
(defn caesarCipher [s k]
  (apply str (map #(char (Integer. %))
                  (map (fn [c]
                         (if (Character/isLetter c)
                           (+ k (int c))
                           (int c))) (char-array s)))))

    ; (def fptr (get (System/getenv) "OUTPUT_PATH"))

    ; (def n (Integer/parseInt (clojure.string/trim (read-line))))


(def s "middle-Outz")

(def k 2)

(println (caesarCipher s k))

    ; (spit fptr (str result "\n") :append true)
