(require '[clojure.string :as str])

(def e2i (slurp "../input/02-e2i.txt"))
(def p2i (slurp "../input/02-p2i.txt"))

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
    (= char "X") loss
    (= char "Y") draw
    (= char "Z") win))

(defn translate
  "translate input values into RPS, part 2"
  [input]
  (let [opponent-choice (RPS (nth input 0)) outcome (RPS (nth input 1))]
    (cond
      (= outcome draw) [opponent-choice opponent-choice]
      (and (= opponent-choice rock) (= outcome loss)) [opponent-choice scissor]
      (and (= opponent-choice rock) (= outcome win)) [opponent-choice paper]
      (and (= opponent-choice paper) (= outcome loss)) [opponent-choice rock]
      (and (= opponent-choice paper) (= outcome win)) [opponent-choice scissor]
      (and (= opponent-choice scissor) (= outcome loss)) [opponent-choice paper]
      (and (= opponent-choice scissor) (= outcome win)) [opponent-choice rock]
      )))

(defn p2s
  ([input]
   (->> input
        str/split-lines             ; newline split
        (map #(str/split % #" "))   ; split on space
        (map #(translate %))        ; turn ascii into values
        (map #(score %))            ; score each round
        (apply +))))                ; sum scores

(println "e2i" (p2s e2i))
(println "p2i" (p2s p2i))
