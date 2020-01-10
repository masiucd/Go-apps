import * as React from 'react';
import { StyledInfo, Text } from './Info.styles';
import Title from '../title/Title';
import { BtnLink } from '../Buttons';

interface P {}

const Info: React.FC<P> = () => {
  return (
    <StyledInfo>
      <Title mainTitle="How about" subTitle="our story" className="info" />
      <Text>
        <p>
          A con panna body cultivar lungo saucer black grounds. Redeye white
          cappuccino strong coffee black froth redeye. Breve at sweet single
          origin, wings kopi-luwak mazagran extra flavour robust. Medium, black
          irish latte mug, cappuccino whipped body galão medium mocha. That,
          carajillo steamed frappuccino pumpkin spice cup rich. Espresso saucer
        </p>
      </Text>
      <BtnLink fade to="/about" bg="black" border="dark">
        About
      </BtnLink>
    </StyledInfo>
  );
};
export default Info;
