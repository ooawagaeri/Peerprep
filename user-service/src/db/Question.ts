/* eslint import/no-cycle: 0 */
import {
  Check,
  Column,
  CreateDateColumn,
  Entity,
  OneToMany,
  PrimaryGeneratedColumn,
  UpdateDateColumn,
} from 'typeorm';
import History from './History';
import { QuestionDifficulty, Question } from '../proto/types';

@Entity('questions')
@Check('difficulty > 0')
export default class QuestionEntity implements Question {
  @PrimaryGeneratedColumn({ name: 'question_id' })
    questionId!: number;

  @Column({ unique: true, nullable: false })
    name!: string;

  @Column({ nullable: false })
    difficulty!: QuestionDifficulty;

  @Column({ nullable: false })
    content!: string;

  @Column({ nullable: false })
    solution!: string;

  @OneToMany(() => History, (history) => history.question)
    histories?: History[];

  @CreateDateColumn({ name: 'create_timestamp' })
    createDateTime?: Date;

  @UpdateDateColumn({ name: 'update_timestamp' })
    updateDateTime?: Date;
}
