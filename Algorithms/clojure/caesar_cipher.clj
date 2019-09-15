(def alphabet-total 26)

(defn adjust-k [c k]
  (let [k (+ k (int c))]
    (if (or (and (Character/isLowerCase c) (> k (int \z)))
            (and (Character/isUpperCase c) (> k (int \Z))))
      (- k alphabet-total)
      k)))

; Complete the caesarCipher function below.

(defn caesarCipher [s k]
  (apply str (map #(char (Integer. %))
                  (map (fn [c]
                         (if (Character/isLetter c)
                           (adjust-k c k)
                           (int c))) (char-array s)))))
                           
(def fptr (get (System/getenv) "OUTPUT_PATH"))

(def n (Integer/parseInt (clojure.string/trim (read-line))))

(def s (read-line))

(def k (Integer/parseInt (clojure.string/trim (read-line))))

(def result (caesarCipher s k))

(spit fptr (str result "\n") :append true)
