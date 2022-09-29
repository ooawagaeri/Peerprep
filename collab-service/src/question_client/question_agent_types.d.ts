import { Question, QuestionDifficulty } from '../proto/question-service';

declare interface IQuestionAgent {
  getQuestionByDifficulty(difficulty: QuestionDifficulty): Promise<Question | undefined>
}