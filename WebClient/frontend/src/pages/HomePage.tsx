import { Stack } from '../styles/primitives/Stack';
import { Text } from '../styles/primitives/Text';
import { Card } from '../styles/component/Card';
import { FeatureItem } from '../styles/component/FeatureItem';
import styled from 'styled-components';
import { Center } from '../styles/layout/Center';
import { Link } from 'react-router-dom';

const CardHome = styled(Card)`
  whith: 60%;
`

const HomePage = () => {
  return (
    <Center>
      <CardHome>
        <Stack gap={16}>
          
          <Text variant="h1">Welcome to Test App</Text>
          <Text variant="muted">
            A platform for passing and managing testing
          </Text>
          
          <Text variant="h2">Platform features</Text>

          <FeatureItem
            title="Online testing"
            description="Take the tests at a convenient time from anywhere"
          />
          <FeatureItem
            title="Results analysis"
            description="View detailed statistics and progress"
          />
          <FeatureItem
            title="User profile"
            description="Manage your data and test history"
          />
          
          <Link to="/login">
            <Text variant='h2'>log in</Text>
          </Link>
        </Stack>
      </CardHome>
    </Center>
  );
};

export default HomePage
