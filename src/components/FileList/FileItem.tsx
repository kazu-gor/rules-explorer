import { basename, dirname } from 'node:path';
import { Badge } from '@inkjs/ui';
import { Box, Text } from 'ink';
import React from 'react';
import { match } from 'ts-pattern';
import type { ClaudeFileInfo } from '../../_types.js';

type FileItemProps = {
  readonly file: ClaudeFileInfo;
  readonly isSelected: boolean;
  readonly isFocused: boolean;
};

export const FileItem = React.memo(function FileItem({
  file,
  isSelected,
  isFocused,
}: FileItemProps): React.JSX.Element {
  // File type badge color and label
  const getFileBadge = (file: ClaudeFileInfo) => {
    return match(file.type)
      .with('claude-md', () => ({ color: 'blue' as const, label: 'PROJECT' }))
      .with('claude-local-md', () => ({
        color: 'yellow' as const,
        label: 'LOCAL',
      }))
      .with('slash-command', () => ({
        color: 'green' as const,
        label: 'COMMAND',
      }))
      .with('global-md', () => ({ color: 'magenta' as const, label: 'GLOBAL' }))
      .with('cursor-rule', () => ({ color: 'gray' as const, label: 'CURSOR' }))
      .with('unknown', () => ({ color: 'gray' as const, label: 'FILE' }))
      .exhaustive();
  };

  // File type icon
  const getFileIcon = (file: ClaudeFileInfo): string => {
    return match(file.type)
      .with('claude-md', () => '📝')
      .with('claude-local-md', () => '🔒')
      .with('slash-command', () => '⚡')
      .with('global-md', () => '🌐')
      .with('cursor-rule', () => '🎯')
      .with('unknown', () => '📄')
      .exhaustive();
  };

  // Get filename and parent directory
  const fileName = basename(file.path);
  const dirPath = dirname(file.path);
  const parentDir = basename(dirPath);

  // Display name (including parent directory)
  // Special handling for home directory
  const displayName =
    file.type === 'global-md'
      ? `~/.claude/${fileName}`
      : file.type === 'slash-command'
        ? fileName.replace('.md', '') // Remove .md for commands
        : file.type === 'cursor-rule'
          ? fileName.replace('.md', '') // Remove .md for cursor rules
          : `${parentDir}/${fileName}`;

  const prefix = isFocused ? '► ' : '  ';

  const fileBadge = getFileBadge(file);

  return (
    <Box justifyContent="space-between" width="100%">
      <Box>
        {isSelected ? (
          <Text backgroundColor="blue" color="white">
            {prefix}
            {getFileIcon(file)} {displayName}
          </Text>
        ) : isFocused ? (
          <Text color="white">
            {prefix}
            {getFileIcon(file)} {displayName}
          </Text>
        ) : (
          <Text>
            {prefix}
            {getFileIcon(file)} {displayName}
          </Text>
        )}
      </Box>
      <Box>
        <Badge color={fileBadge.color}>{fileBadge.label}</Badge>
      </Box>
    </Box>
  );
});
