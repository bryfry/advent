(def e1i (slurp "../input/02-e1i.txt"))
(def p1i (slurp "../input/02-p1i.txt"))
;(def e2i (slurp "../input/01-e2i.txt"))
;(def p2i (slurp "../input/01-p2i.txt"))

(def loss 0)
(def draw 3)
(def win 6)
(def rock 1)
(def paper 2)
(def scissor 3)


(defn score
  "Score for user given RPS input"
  [opponent-choice user-choice]
  (+ user-choice
    (cond
      (= opponent-choice user-choice) draw
      (and (= opponent-choice rock) (= user-choice scissor)) loss
      (and (= opponent-choice rock) (= user-choice paper)) win
      (and (= opponent-choice paper) (= user-choice rock)) loss
      (and (= opponent-choice paper) (= user-choice scissor)) win
      (and (= opponent-choice scissor) (= user-choice paper)) loss
      (and (= opponent-choice scissor) (= user-choice rock)) win )))

(defn p1s
  ([input]
   (->> input
        str/split-lines             ; newline split
        )))

(println (p1s e1i))
;(println (score rock paper)
;(score paper rock)
;(score scissor scissor))
