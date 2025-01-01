"use client";
import { useState } from "react";
import style from "./carousel.module.css";
import forwardArrow from "@/../public/svg/Arrow_forward.svg";
import backArrow from "@/../public/svg/Arrow_back.svg";
import Image from "next/image";

export default function Carousel({ images }: { images: string[] }) {
  const [imageIndex, setImageIndex] = useState(() => 0);

  return (
    <div className={style.carousel}>
      <div className={style.content}>
        <img src={images[imageIndex]} alt={""} className={style.image}></img>
      </div>

      <div className={style.buttons}>
        {imageIndex !== 0 && (
          <button
            onClick={() => setImageIndex(imageIndex - 1)}
            className={`${style["back-button"]} ${style.button}`}
          >
            <Image src={backArrow} alt="" />
          </button>
        )}
        {imageIndex !== images.length - 1 && (
          <button
            onClick={() => setImageIndex(imageIndex + 1)}
            className={`${style["forward-button"]} ${style.button}`}
          >
            <Image src={forwardArrow} alt="" />
          </button>
        )}
      </div>
    </div>
  );
}
