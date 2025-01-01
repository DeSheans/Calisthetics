import { Program, Training as TrainingType } from "@/interfaces/interfaces";
import styles from "./page.module.css";

export default async function ProgramPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const id = (await params).id;
  const apiUrl = process.env.NEXT_PRIVATE_API_URL;
  const data: Promise<Program> = fetch(`http://${apiUrl}/programs/${id}`).then(
    (res) => res.json()
  );
  if (data === null) return <>ERROR 404</>;

  return (
    <div className={styles.inner}>
      <h1 className={styles.heading}>{(await data).name}</h1>
      <p className={styles.description}>{(await data).description}</p>

      <div className={styles.brief}>
        <div className={styles.briefInner}>
          <span>{`Тип программы: ${(await data).programType}`}</span>
          <span>{`Сложность: ${(await data).difficulty}`}</span>
          {/* <button>В избранное</button> */}
        </div>
      </div>

      <div className={styles.trainingsContainer}>
        {(await data).trainings.map((t) => (
          <Training key={t.id} training={t}></Training>
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
          <tr key={e.exercise} className={styles.trainingRow}>
            <td className={styles.trainingData}>{e.exercise}</td>
            <td className={styles.trainingData}>{e.sets}</td>
            <td className={styles.trainingData}>{e.reps}</td>
            <td className={styles.trainingData}>{e.rest}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}
