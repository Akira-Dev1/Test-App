import { Stack } from '../primitives/Stack';
import { Text } from '../primitives/Text';
import { Input } from '../primitives/Input';

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export function FormField({ label, ...props }: any) {
  return (
    <Stack gap={6}>
      <Text variant="muted">{label}</Text>
      <Input {...props} />
    </Stack>
  );
}