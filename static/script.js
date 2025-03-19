const questions = [
    ["test1", true],
    ["test2", false],
    ["test3", true],
    ["test4", false],
    ["test5", true]
];

let score = 0;
let currentQuestionIndex = 0;

const questionElement = document.getElementById("question");
const resultElement = document.getElementById("result");
const scoreElement = document.getElementById("score");

const trueButton = document.getElementById("trueButton");
const falseButton = document.getElementById("falseButton");

function askQuestion() {
    if (currentQuestionIndex < questions.length) {
        const question = questions[currentQuestionIndex][0];
        questionElement.textContent = question;
        resultElement.textContent = "";
    } else {
        questionElement.textContent = "Jeu terminé !";
        resultElement.textContent = `Votre score est de ${score}/${questions.length}.`;
        trueButton.disabled = true;
        falseButton.disabled = true;
    }
}

function checkAnswer(userAnswer) {
    const correctAnswer = questions[currentQuestionIndex][1];
    if (userAnswer === correctAnswer) {
        resultElement.textContent = "Bonne réponse !";
        score++;
    } else {
        resultElement.textContent = "Mauvaise réponse !";
    }
    currentQuestionIndex++;
    scoreElement.textContent = `Score: ${score}`;
    setTimeout(askQuestion, 1000); // Attendre avant de poser la question suivante
}

trueButton.addEventListener("click", () => checkAnswer(true));
falseButton.addEventListener("click", () => checkAnswer(false));

// Initialiser le jeu
askQuestion();
