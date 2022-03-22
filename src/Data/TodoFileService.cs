namespace ChaosTasks.Services;

using System.IO;
using System.Collections.Immutable;
using ChaosTasks.Models;

public class TodoFileService : IDisposable
{
  private FileSystemWatcher _watcher;
  private bool disposed = false;

  public void ListenForChanges(string absolutePath, string fileName, FileSystemEventHandler changeEvent) {
    _watcher = new FileSystemWatcher(absolutePath);
    _watcher.NotifyFilter = NotifyFilters.LastWrite;
    _watcher.Changed += changeEvent;
    _watcher.Filter = fileName;
    _watcher.EnableRaisingEvents = true;
  }

  public void Dispose() {
      Dispose(disposing: true);
      GC.SuppressFinalize(this);
  }

  protected virtual void Dispose(bool disposing) {
    if (!this.disposed) {
      if (disposing) {
	this._watcher.Dispose();
      }
    }
  }
}
