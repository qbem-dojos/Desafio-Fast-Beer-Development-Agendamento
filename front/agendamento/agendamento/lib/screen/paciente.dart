import 'package:flutter/material.dart';
import 'package:syncfusion_flutter_calendar/calendar.dart';


class PacienteScreen extends StatefulWidget {
  const PacienteScreen({super.key});

  @override
  State<PacienteScreen> createState() => _PacienteScreenState();
}

class _PacienteScreenState extends State<PacienteScreen> {
  @override
  Widget build(BuildContext context) {
    return SfCalendar(
        view: CalendarView.workWeek,
        dataSource: _getCalendarDataSource(),
        onTap: (calendarTapDetails) => popFormulario(calendarTapDetails));
  }

  Future popFormulario(CalendarTapDetails calendarTapDetails) {
    if (calendarTapDetails.appointments != null) {
      return showDialog<String>(
        context: context,
        builder: (BuildContext context) => AlertDialog(
          title: const Text('Agendamento'),
          content: const Text('Essa data não está disponível'),
          actions: <Widget>[
            TextButton(
              onPressed: () => Navigator.pop(context, 'Cancel'),
              child: const Text('Cancel'),
            ),
            TextButton(
              onPressed: () => Navigator.pop(context, 'OK'),
              child: const Text('OK'),
            ),
          ],
        ),
      );
    }
    return showDialog<String>(
      context: context,
      builder: (BuildContext context) => AlertDialog(
        title: const Text('Agendamento'),
        content: const Text('Preencha os seus dados para agendar sua consulta'),
        actions: formulario(context),
      ),
    );
  }

  List<Widget> formulario(BuildContext context) {
    return <Widget>[
      TextFormField(
        decoration: const InputDecoration(
          hintText: 'E-mail',
        ),
        validator: (String? value) {
          if (value == null || value.isEmpty) {
            return 'Please enter some text';
          }
          return null;
        },
      ),
      TextFormField(
        decoration: const InputDecoration(
          hintText: 'Nome',
        ),
        validator: (String? value) {
          if (value == null || value.isEmpty) {
            return 'Please enter some text';
          }
          return null;
        },
      ),
      TextFormField(
        decoration: const InputDecoration(
          hintText: 'Telefone',
        ),
        validator: (String? value) {
          if (value == null || value.isEmpty) {
            return 'Please enter some text';
          }
          return null;
        },
      ),
      TextButton(
        onPressed: () => Navigator.pop(context, 'Cancel'),
        child: const Text('Cancel'),
      ),
      TextButton(
        onPressed: () => Navigator.pop(context, 'OK'),
        child: const Text('OK'),
      ),
    ];
  }
}

class DataSource extends CalendarDataSource {
  DataSource(List<Appointment> source) {
    appointments = source;
  }
}

DataSource _getCalendarDataSource() {
  List<Appointment> appointments = <Appointment>[];
  appointments.add(Appointment(
      startTime: DateTime.now(),
      endTime: DateTime.now().add(const Duration(hours: 2)),
      isAllDay: false,
      subject: 'Meeting',
      color: Colors.blue,
      startTimeZone: 'America/Sao_Paulo',
      endTimeZone: 'America/Sao_Paulo'));
  return DataSource(appointments);
}
