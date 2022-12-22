(require '[clojure.string :as str])

(def e1i (slurp "../input/01-e1i.txt"))
(def p1i (slurp "../input/01-p1i.txt"))
(def e2i (slurp "../input/01-e2i.txt"))
(def p2i (slurp "../input/01-p2i.txt"))


(defn p1s
  ([input]
   (->> input
        str/split-lines             ; newline split
        (map parse-long)            ; each item an int (empty line = nil)
        (partition-by nil?)         ; into sub lists where nil is deliminator
        (filter #(not-any? nil? %)) ; get rid of the nil sub-lists
        (map #(apply + %))          ; sum
        (apply max))))              ; max

(defn p2s
  ([input]
   (->> input
        str/split-lines             ; newline split
        (map parse-long)            ; each item an int (empty line = nil)
        (partition-by nil?)         ; into sub lists where nil is deliminator
        (filter #(not-any? nil? %)) ; get rid of the nil sub-lists
        (map #(apply + %))          ; sum
        (sort)                      ; sort assending
        (take-last 3)               ; top 3
        (apply +))))                ; sum top 3

(println "e1s" (p1s e1i))
(println "p1s" (p1s p1i))
(println "e2s" (p2s e2i))
(println "p2s" (p2s p2i))
