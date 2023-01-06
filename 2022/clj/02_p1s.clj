(require '[clojure.string :as str])

(def e1i (slurp "../input/02-e1i.txt"))
(def p1i (slurp "../input/02-p1i.txt"))

(def loss 0)
(def draw 3)
(def win 6)
(def rock 1)
(def paper 2)
(def scissor 3)

(defn score
  "Score for user given RPS input"
  [input]
  (let [opponent-choice (nth input 0) user-choice (nth input 1)]
  (+ user-choice
    (cond
      (= opponent-choice user-choice) draw
      (and (= opponent-choice rock) (= user-choice scissor)) loss
      (and (= opponent-choice rock) (= user-choice paper)) win
      (and (= opponent-choice paper) (= user-choice rock)) loss
      (and (= opponent-choice paper) (= user-choice scissor)) win
      (and (= opponent-choice scissor) (= user-choice paper)) loss
      (and (= opponent-choice scissor) (= user-choice rock)) win ))))

(defn RPS
  [char]
  (cond
    (= char "A") rock
    (= char "B") paper
    (= char "C") scissor
    (= char "X") rock
    (= char "Y") paper
    (= char "Z") scissor))

(defn translate
  "translate input values into RPS, part 1"
  [coll]
  (map RPS coll))

(defn p1s
  ([input]
   (->> input
        str/split-lines             ; newline split
        (map #(str/split % #" "))   ; split on space
        (map #(translate %))        ; turn ascii into RPS values
        (map #(score %))            ; score each round
        (apply +))))                ; sum scores


(println "e1i" (p1s e1i))
(println "p1i" (p1s p1i))
