import { Program, Training as TrainingType } from "@/interfaces/interfaces";
import styles from "./page.module.css";
import { twoDaySplit } from "@/stubs/programStub";

const data: Program = twoDaySplit;

export default function ProgramPage() {
  return (
    <div className={styles.inner}>
      <h1 className={styles.heading}>{data.name}</h1>
      <p className={styles.description}>{data.description}</p>

      <div className={styles.brief}>
        <div className={styles.briefInner}>
          <span>{`Тип программы: ${data.type}`}</span>
          <span>{`Сложность: ${data.difficulty}`}</span>
          {/* <button>В избранное</button> */}
        </div>
      </div>

      <div className={styles.trainingsContainer}>
        {data.trainings.map((t) => (
          <Training training={t}></Training>
        ))}
      </div>
    </div>
  );
}

function Training({ training }: { training: TrainingType }) {
  return (
    <table className={styles.trainingTable}>
      <caption className={styles.trainingCaption}>{training.name}</caption>
      <thead>
        <tr className={`${styles.trainingRow}`}>
          <th className={styles.trainingHeading}>Упражнение</th>
          <th className={styles.trainingHeading}>Подходы</th>
          <th className={styles.trainingHeading}>Повторения</th>
          <th className={styles.trainingHeading}>Перерыв между подходами</th>
        </tr>
      </thead>
      <tbody>
        {training.exercises.map((e) => (
          <tr className={styles.trainingRow}>
            <td className={styles.trainingData}>{e.exercise.name}</td>
            <td className={styles.trainingData}>{e.sets}</td>
            <td className={styles.trainingData}>{e.reps}</td>
            <td className={styles.trainingData}>{e.rest}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}
