import { Exercise } from "@/interfaces/interfaces";
import styles from "./exerciseBrief.module.css";

export default function ExerciseBrief({
  exercise,
  isFavorite,
}: {
  exercise: Exercise;
  isFavorite: boolean;
}) {
  return (
    <div className={styles.brief}>
      <div className={styles.briefInner}>
        <div className={styles.briefExerciseType}>
          {`Тип упражнения: ${exercise.exerciseType}`}
        </div>
        <div className={styles.briefEquipment}>
          <span>Требуемое оборудование: </span>
          {exercise.equipment.length === 0 ? (
            "нет"
          ) : (
            <ul className={styles.briefList}>
              {exercise.equipment.map((e) => (
                <li>{e}</li>
              ))}
            </ul>
          )}
        </div>
        <div className={styles.briefDifficulty}>
          {`Сложность: ${exercise.difficulty}`}
        </div>
        <div className={styles.briefMuscles}>
          <span>Целевые группы мышц: </span>
          <ul className={styles.briefList}>
            {exercise.muscles.map((m) => (
              <li>{m}</li>
            ))}
          </ul>
        </div>

        {/* <button className={styles.briefFavoriteButton}>Избранное</button> */}
      </div>
    </div>
  );
}
