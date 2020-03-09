enum StepMoveDirection { Up, Down }

extension StepMoveDirectionExtension on StepMoveDirection {
  String get queryString {
    switch (this) {
      case StepMoveDirection.Up:
        return "up";
      case StepMoveDirection.Down:
        return "down";
    }
    return "unknown direction";
  }
}
