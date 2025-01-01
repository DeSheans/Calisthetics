import type { ExerciseCard } from "@/interfaces/interfaces";
import styles from "./exerciseCard.module.css";
import Link from "next/link.js";

export default function ExerciseCard({ exercise }: { exercise: ExerciseCard }) {
  let exerciseUrl = `exercises/${exercise.id.toString()}`;

  return (
    <div className={styles.card}>
      <Link className={styles.link} href={exerciseUrl}>
        <img
          className={styles.picture}
          src={`/pictures/${exercise.picture}`}
          alt={exercise.name}
          loading="lazy"
          decoding="async"
        />
      </Link>

      <div className={styles.bot}>
        <Link className={styles.link} href={exerciseUrl}>
          <div className={styles.name}>{exercise.name}</div>
        </Link>

        <div className={styles.class}>
          {`${exercise.exerciseType}, сложность: ${exercise.difficulty}`}
        </div>
        <div className={styles.muscles}>
          <ul>
            {exercise.muscles.slice(0, 3).map((m) => (
              <li key={m}>{m}</li>
            ))}
          </ul>
        </div>
      </div>
    </div>
  );
}
