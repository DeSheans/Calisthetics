import type { ProgramCard, Training } from "@/interfaces/interfaces";
import Favorite from "@/components/favorite/favoriteSvg/favorite";
import styles from "./programCard.module.css";
import Link from "next/link";

export default function Card({ Program }: { Program: ProgramCard }) {
  return (
    <div className={styles.card}>
      <div className={styles.inner}>
        <div className={styles.heading}>
          <Link className={styles.link} href={`programs/${Program.id}`}>
            {Program.name}
          </Link>
        </div>
        <div className={styles.class}>
          <span className={styles.type}>Тип: {Program.programType}. </span>
          <span className={styles.trainingNumber}>
            Количество тренировок: {Program.trainings.length}.{" "}
          </span>
          <span className={styles.difficulty}>
            Сложность: {Program.difficulty}.
          </span>
        </div>
        <ul className={styles.trainings}>
          {Program.trainings.map((t) => (
            <li key={t.id}>
              <Training key={t.id} training={t}></Training>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}

function Training({ training }: { training: Training }) {
  return (
    <div className={styles.training}>
      <span>{training.name}</span>
      <span> | </span>
      <span>{training.exercises.map((e) => e.exercise).join("; ")}</span>
    </div>
  );
}
