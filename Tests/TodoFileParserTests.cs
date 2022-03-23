namespace ChaosTasks.Tests;

using System;
using System.IO;
using NUnit.Framework;
using ChaosTasks.Data;

public class TodoFileParserTests
{
  private TodoParserService parser;
  private string[] lines;

  [SetUp]
  public void Setup()
  {
    this.parser = new TodoParserService();
    this.lines = File.ReadAllLines(@"/Users/acbilson/source/chaos-tasks/Tests/data/basic.txt");
  }

  [TearDown]
  public void TearDown()
  {
    this.parser = null;
    this.lines = null;
  }

  [Test]
  public void LoadsData()
  {
    var todoList = this.parser.ToTodoList(this.lines);
    Assert.Greater(todoList.Count, 0);
    Console.WriteLine("Testing raw strings");
    todoList.ForEach(todo => Console.WriteLine(todo.Raw));
    todoList.ForEach(todo => Console.WriteLine(todo.Text));
    todoList.ForEach(todo => Console.WriteLine(todo.Priority));
    todoList.ForEach(todo => Console.WriteLine(string.Join(',', todo.Projects)));
    todoList.ForEach(todo => Console.WriteLine(string.Join(',', todo.Contexts)));
  }
}
