import { Exercise } from "@/interfaces/interfaces";
import { easy } from "@/stubs/difficultyStub";
import { push } from "@/stubs/exerciseTypeStub";
import { deltoids, pectoral, triceps } from "@/stubs/muscleStubs";
import styles from "./page.module.css";
import ExerciseBrief from "@/components/exerciseBrief/exerciseBrief";

const data: Exercise = {
  id: 0,
  name: "Отжимание",
  muscles: [triceps, deltoids, pectoral],
  difficulty: easy,
  exerciseType: push,
  equipment: [],
  description:
    "Lorem ipsum dolor sit amet consectetur adipisicing elit. Deleniti, laudantium obcaecati. Impedit odio placeat dolorum ipsa sed repellendus dolor, est totam tenetur, amet molestias dolorem necessitatibus dolores culpa soluta minima?",
  tips: [
    "Molestias dolorem necessitatibus dolores culpa soluta minima?",
    "Adipisicing elit. Deleniti, ipsa sed repellendus dolor, est totam tenetur, dolorem necessitatibus.",
  ],
  variations: [
    "Amet molestias dolorem necessitatibus dolores culpa soluta minima?",
    "Consectetur adipisicing elit. Deleniti, dio placeat dolorum ipsa sed repellendus dolor, est totam tenetur.",
    "Lorem met molestias dolorem necessitatibus dolores culpa soluta minima?",
  ],
  pictures: [],
};

// const isFavorite = false;
const similar = true;

export default function ExercisePage() {
  return (
    <div className={styles.inner}>
      <h1 className={styles.heading}>{data.name}</h1>
      <div className={styles.carousel}></div>

      <div className={styles.briefContainer}>
        <ExerciseBrief exercise={data} isFavorite={true}></ExerciseBrief>
      </div>

      <div className={styles.descriptionContainer}>
        <Description
          header="Техника выполнения"
          text={data.description}
        ></Description>
        <List header="Советы:" listItems={data.tips}></List>
        <List header="Варианты усложнения:" listItems={data.variations}></List>
      </div>

      {similar && (
        <div className={styles.similar}>
          <hr className={styles.hr}></hr>
          <h2>Похожие упражнения</h2>
          <div className={styles.similarItemsContainer}>
            {[0, 0, 0, 0, 0].map(() => (
              <div className={styles.similarItem}></div>
            ))}
          </div>
        </div>
      )}
    </div>
  );
}

function Description({ header, text }: { header: string; text: string }) {
  return (
    text && (
      <div className={styles.description}>
        <h2>{header}</h2>
        <p>{text}</p>
      </div>
    )
  );
}

function List({ header, listItems }: { header: string; listItems: string[] }) {
  let c = 0;
  return (
    listItems && (
      <div className={styles.list}>
        <h2>{header}</h2>
        {listItems.map((i) => (
          <li key={c++}>{i}</li>
        ))}
      </div>
    )
  );
}
