import * as React from 'react';
import { StyledIntro, Text } from './Intro.styles';
import Title from '../../styled/title/Title';

interface P {}

const Intro: React.FC<P> = () => {
  return (
    <StyledIntro>
      <Title mainTitle="Worlds Best" subTitle="story" flex />
      <Text>
        <p>
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Ab culpa
          ullam, magnam perferendis illo, vitae velit unde vero quas illum
          repellat sed quod nam, deserunt distinctio ipsa ipsam corrupti ut.
          ullam, magnam perferendis illo, vitae velit unde vero quas illum
          repellat sed quod nam, deserunt distinctio ipsa ipsam corrupti ut.
          ullam, magnam perferendis illo, vitae velit unde vero quas illum
          repellat sed quod nam, deserunt distinctio ipsa ipsam corrupti ut.
        </p>
      </Text>
    </StyledIntro>
  );
};
export default Intro;
