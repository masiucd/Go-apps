import * as React from 'react';
import { StyledFoodGrid, GridItem } from './Styles.foodgrid';
import { useStaticQuery, graphql } from 'gatsby';
import Img, { FluidObject } from 'gatsby-image';

const queryData = graphql`
  query {
    img: file(relativePath: { eq: "homeGallery/img-1.jpeg" }) {
      childImageSharp {
        fluid(maxWidth: 500) {
          ...GatsbyImageSharpFluid_tracedSVG
          # ...GatsbyImageSharpFluid_withWebp_tracedSVG
        }
      }
    }
  }
`;

interface Props {
  img: {
    childImageSharp: {
      fluid?: FluidObject | FluidObject[] | undefined;
    };
  };
}

const FoodGrid: React.FC<Props> = () => {
  const data: Props = useStaticQuery(queryData);
  // console.log('img', data);
  return (
    <StyledFoodGrid>
      <GridItem className="main">
        <Img fluid={data.img.childImageSharp.fluid} />
      </GridItem>
    </StyledFoodGrid>
  );
};
export default FoodGrid;
